#!/bin/bash

# This is a Cross-Compiling build script for the data manager microservice
# The build script generates linux binaries
# Usage:
# build.sh [-v] [command] [targetOS]
#
VERSION="1.0.0"

echo "Build Version: $VERSION"
echo ""

case "$OSTYPE" in
	darwin*)
		export HOST_OS="darwin"
		;;
	linux*)
		export HOST_OS="linux"
		;;
	*)
		echo "unknown OS Type: $OSTYPE"
		exit 1
		;;
esac

MS_NAME="data_manager"
PKG_LIST="models stats template"
REPO_PATH="${GOPATH}/src/github.com/vmtrain/data_manager"
BUILD_IMAGE="docker.io/golang:1.6"
GO_VERSION="1.6"
VERBOSE=0
VERBOSEFLAG=>/dev/null
COVER_FILE=testcoverage.out
CANTRAN=false


function showTestCoverage {
	# Now lets see how much test coverage we have for each packge
	display ""
	display "================================================================================"
	echo "Calculating Test Coverage:"
	display "================================================================================"

	# for each pakcage in the package list - we change to that directory to run the
	# test coverage tools.    go test -coverprofile does not support multiple packages at once

    echo -e "\tCommands:"
    for cmd in $MS_NAME; do
        pushd "${REPO_PATH}/cmd/${cmd}" > /dev/null

        rm -f $COVER_FILE
        go test -coverprofile $COVER_FILE > /dev/null 2>&1
        if [ -f $COVER_FILE ]
        then
            OUTPUT="$(go tool cover -func=$COVER_FILE | grep total:)"
            COLS=( $OUTPUT )
            echo -e "\t\t${cmd}: ${COLS[2]}"
            rm -f $COVER_FILE
        fi
        popd > /dev/null
    done;

    echo -e "\tPackages:"
	for pkg in $PKG_LIST; do
		pushd "${REPO_PATH}/pkg/${pkg}" > /dev/null

		rm -f $COVER_FILE
		go test -coverprofile $COVER_FILE > /dev/null 2>&1
        if [ -f $COVER_FILE ]
        then
            OUTPUT="$(go tool cover -func=$COVER_FILE | grep total:)"
            COLS=( $OUTPUT )
            echo -e "\t\t${pkg}: ${COLS[2]}"
            rm -f $COVER_FILE
        fi
		popd > /dev/null
	done;
}

function testApps() {
	display "================================================================================"
	echo "Running Tests - fmt, vet, lint, test"
	display "================================================================================"


	echo -n "FMTing Everything"
	fmt=$(go fmt ./... | grep -v buildinfo | xargs)
        if [[ ($? != 0) || ($fmt != "") ]]; then
            echo " - Error!"
            echo $fmt
            echo "go fmt ./... failed.  some files were not formatted correctly"
            exit 1
        fi;
        echo " - OK"

	echo -n "Vetting Everything"
	output=$(go vet ./... 2>&1)
	if [[ $? != 0 ]]; then
                echo " - Error!"
		echo $output
		exit 1
	fi;
        echo " - OK"

	# And we will do lint
	# We may need to go get lint first
	if [ "$TARGET_OS" == "$HOST_OS" ]; then
		if ! type "golint" > /dev/null 2>&1 ; then
			echo "Getting golint Linter"
			go get -u  "github.com/golang/lint/golint"
		fi
	fi

	# older version of lint do not support the set_exit_status flag, so we check the output as well
	echo -n "Linting Everything"
	output=$(golint -set_exit_status ./... 2>&1)
	if [[($? != 0) || ("$output" != "") ]]; then
                echo " - Error!"
		echo $output
		echo "Lint Failed"
		exit 1
	fi;
        echo " - OK"


	rm -f gotest.out

	# We do testing in the geronimo directory to ensure we are testing all the packages
	# and not just the apps
	# We call canticle to resolve dependencies

	echo -n "Running tests"
	go test -v  ./...  >>  gotest.out 2>&1
	if [[ $? != 0 ]]; then
                echo " - Error!"
		cat gotest.out | sed ''/FAIL:/s//$(printf "\033[31mFAIL:\033[0m")/''
		echo ""
		echo "================================================================================"
		echo "FAIL Summary:"
		echo "================================================================================"
		grep "FAIL" gotest.out
		echo "!!! Testing failed. !!!"
		exit 1
	fi;
        echo " - OK"

	if [[ ("$VERBOSE" = 1) ]]; then
		cat gotest.out
		echo ""
		echo "Tests succeeded"
	fi



	showTestCoverage
}


makeDocker() {
	display ""
	display "================================================================================"
	echo "Deploying Docker image:"
	display "================================================================================"
	if ! type "docker" > /dev/null 2>&1; then
		# Can't find docker ;)    Let's install it
		echo "ERROR: Cannot run docker"
		exit 1
	fi
	if [ "$HOST_OS" == "darwin" ]; then
		eval $(docker-machine env dev-dev)
	fi
    docker build -t $1 .
    if [ $? -ne 0 ]; then 
        echo "error building docker image"
        exit 1
    fi
}

function build() {
    # we should already have the dependencies from canticle
	# now let's build it
	echo "Building $1"
	go build $VERBOSEFLAG
}

function buildApps() {
    display "================================================================================"
    echo "Building Apps"
    display "================================================================================"

	for cmd in $MS_NAME; do
	    display "================================================================================"
	    display "BUILDING $cmd PACKAGE"
	    display "================================================================================"

	    pushd "${REPO_PATH}/cmd/${cmd}" > /dev/null
	    build $cmd
	    if [[ $? != 0 ]]; then
	        echo $cmd build failed.
	        exit 1
	    fi;
	    popd > /dev/null
	done;
}

function cleanApps() {
	display "================================================================================"
	echo "Cleaning Apps"
	display "================================================================================"

	for cmd in $MS_NAME; do
		display "================================================================================"
		display "Cleaning $cmd PACKAGES"
		display "================================================================================"
		pushd "${REPO_PATH}/cmd/${cmd}" > /dev/null
		go clean
		if [[ $? != 0 ]]; then
			echo $cmd clean failed.
			exit 1
		fi;
		popd > /dev/null
	done;
	if [  -f gotest.out ]; then
		echo "Removing gotest.out"
		rm gotest.out
	fi
	if [  -f junit.xml ]; then
		echo "Removing junit.xml"
		rm junit.xml
	fi
}

function SetGoEnvironment {
	# Export the GO Cross Compiling Environment variable
	display "Setting Go Environment"
	display "Build Host OS: $HOST_OS"
	display "Build Target OS: $TARGET_OS"
	# as of go 1.6.1 - cross compiling cgo is not complete
	# Build static linked binary due to issue with Alpine containers using musl instead of glibc
	export CGO_ENABLED=0

	#  Let's set proper Cross Compile Options for our build tools
	case "$HOST_OS" in
		darwin)
			export CC="clang"
			export GOGCCFLAGS="-fPIC -m64 -pthread -fno-caret-diagnostics -Qunused-arguments -fmessage-length=0 -fno-common"
			export CXX="clang++"
			;;
		linux)
			export CC="gcc"
			export GOGCCFLAGS="-fPIC -m64 -fmessage-length=0"
			export CXX="g++"
			;;
		*)
			echo "unknown Platform Type: $HOST_OS"
			exit 1
			;;
	esac
	#  Let's set Target Compile Options
	case "$TARGET_OS" in
		darwin)
			export GOARCH="amd64"
			export GOOS="darwin"
			;;
		linux)
			export GOARCH="amd64"
			export GOOS="linux"
			;;
		*)
			echo "Unknown Target OS: $TARGET_OS"
			exit 1
			;;
	esac
	if [ "$VERBOSE" = 1 ]; then
		echo "GO Environment"
		go version
		echo ""
		echo "PATH = $PATH"
		echo ""
		go env
		echo ""
	fi
}

function display {
	if [ "$VERBOSE" = 1 ]; then
		echo "$1"
	fi
}

function createContainer {
	CGO_ENABLED=0 
	go build -a --installsuffix cgo .
	docker build -t q3-training-journal --rm=true .
}

while getopts ":vc" opt; do
	case $opt in
		v)
			VERBOSE=1
			VERBOSEFLAG="-v"
			;;
		c)
		    # Skip Canticle for more rapid iterations
			CANTRAN=true
			;;

		\?)
			echo "Invalid option: -$OPTARG" >&2
			;;
	esac
done
shift $((OPTIND-1))
COMMAND=${1:-build}
TARGET_OS=$2

# Check version of Go Language
if ! (go version 2> /dev/null | grep "go$GO_VERSION" > /dev/null) ; then
	echo "Expecting Go Version $GO_VERSION.   Current version is: [$(go version)]"
	exit 0
fi
# Make sure GOPATH environment variable is set
if [ -z ${GOPATH+x} ]; then
	echo "GOPATH is not set";
	exit 1
else
	export PATH=$PATH:$GOPATH/bin
fi

if [ -z "$TARGET_OS" ]; then
	# Set target to current OS
	TARGET_OS="$(go env GOOS)"
fi
display "COMMAND $COMMAND, TARGET_OS $TARGET_OS"

# Display our Go Environment info
display ""
display "Build Host OS: $HOST_OS"
display "Build Target OS: $TARGET_OS"

case "$COMMAND" in
	build)
		SetGoEnvironment
		buildApps
		echo "Done"
		;;
	clean)
		cleanApps
		echo "Done"
		;;
	test)
		SetGoEnvironment
		testApps
		echo "Done"
		;;
    containerize)
        echo "Containerizing Apps"
        TARGET_OS='linux'
        SetGoEnvironment
        buildApps
        makeDocker datamanager:$VERSION
        echo created container datamanager:$VERSION
        echo "Done"
        ;;
	
	*)
		echo "USAGE:   build.sh [build|clean|test|containerize]"
		;;
esac