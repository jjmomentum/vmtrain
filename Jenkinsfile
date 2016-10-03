 node ('master'){
  stage 'Build and Validate'
  checkout scm
  sh 'npm install'
  sh 'npm test'
  sh 'echo $JOB_NAME'

  stage 'REST Service Tests'
  def teamId = env.JOB_NAME.split("-")[1].substring(4)
  def port = "31"+teamId
  echo "Running build for team " + teamId
  try {
    sh 'PORT='+port+' npm start &'
    sh 'sleep 5'
    sh '/usr/local/share/SoapUI-5.2.1/bin/testrunner.sh -PTEAM='+teamId+' -PPORT='+port+' ./Q3-Training-Tests-soapui-project.xml'
  } finally {
    stage 'Shutdown and Cleanup'
    sh 'npm stop'
  }
 }


