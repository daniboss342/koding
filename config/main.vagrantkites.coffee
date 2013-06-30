fs              = require 'fs'
nodePath        = require 'path'
deepFreeze      = require 'koding-deep-freeze'

version         = "0.0.1"
mongo           = 'local.koding.com:27017/koding'
projectRoot     = nodePath.join __dirname, '..'
socialQueueName = "koding-social-vagrant"

module.exports =
  aws           :
    key         : 'AKIAJSUVKX6PD254UGAA'
    secret      : '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'
  uri           :
    address     : "http://local.koding.com:80"
  userSitesDomain: 'local.koding.com'
  containerSubnet: "10.128.2.0/9"
  projectRoot   : projectRoot
  version       : version
  webserver     :
    login       : 'prod-webserver'
    port        : 80
    clusterSize : 1
    queueName   : socialQueueName+'web'
    watch       : yes
  sourceServer  :
    enabled     : yes
    port        : 3526
  mongo         : mongo
  neo4j         :
    read        : "http://local.koding.com"
    write       : "http://local.koding.com"
    port        : 7474
  runNeo4jFeeder: yes
  runGoBroker   : yes
  runKontrol    : no
  runRerouting  : yes
  runUserPresence: yes
  runPersistence: yes
  compileGo     : yes
  buildClient   : yes
  runOsKite     : yes
  runProxy      : yes
  misc          :
    claimGlobalNamesForUsers: no
    updateAllSlugs : no
    debugConnectionErrors: yes
  uploads       :
    enableStreamingUploads: no
    distribution: 'https://d2mehr5c6bceom.cloudfront.net'
    s3          :
      awsAccountId        : '616271189586'
      awsAccessKeyId      : 'AKIAJO74E23N33AFRGAQ'
      awsSecretAccessKey  : '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'
      bucket              : 'koding-uploads'
  # loadBalancer  :
  #   port        : 3000
  #   heartbeat   : 5000
    # httpRedirect:
    #   port      : 80 # don't forget port 80 requires sudo
  bitly :
    username  : "kodingen"
    apiKey    : "R_677549f555489f455f7ff77496446ffa"
  authWorker    :
    login       : 'prod-auth-worker'
    queueName   : socialQueueName+'auth'
    numberOfWorkers: 1
    watch       : yes
  social        :
    login       : 'prod-social'
    numberOfWorkers: 1
    watch       : yes
    queueName   : socialQueueName
  cacheWorker   :
    login       : 'prod-social'
    watch       : yes
    queueName   : socialQueueName+'cache'
    run         : no
  graphFeederWorker:
    numberOfWorkers: 2
  presence      :
    exchange    : 'services-presence'
  client        :
    version     : version
    watch       : yes
    watchDuration : 300
    includesPath: 'client'
    websitePath : 'website'
    js          : "js/kd.#{version}.js"
    css         : "css/kd.#{version}.css"
    indexMaster : "index-master.html"
    index       : "default.html"
    useStaticFileServer: no
    staticFilesBaseUrl: 'http://local.koding.com'
    runtimeOptions:
      userSitesDomain: 'local.koding.com'
      useNeo4j: yes
      logToExternal: no  # rollbar, mixpanel etc.
      resourceName: socialQueueName
      suppressLogs: no
      broker    :
        sockJS  : 'http://local.koding.com:8008/subscribe'
      apiUri    : 'https://dev-api.koding.com'
      # Is this correct?
      version   : version
      mainUri   : 'http://local.koding.com'
      appsUri   : 'https://koding-apps.s3.amazonaws.com'
      sourceUri : 'http://local.koding.com:3526'
  mq            :
    host        : 'local.koding.com'
    port        : 5672
    apiAddress  : "local.koding.com"
    apiPort     : 15672
    login       : 'PROD-k5it50s4676pO9O'
    componentUser: "PROD-k5it50s4676pO9O"
    password    : 'djfjfhgh4455__5'
    heartbeat   : 10
    vhost       : '/'
  broker        :
    ip          : ""
    port        : 8008
    certFile    : ""
    keyFile     : ""
  kites:
    disconnectTimeout: 3e3
    vhost       : 'kite'
  email         :
    host        : 'local.koding.com'
    protocol    : 'http:'
    defaultFromAddress: 'hello@koding.com'
  emailWorker   :
    cronInstant : '*/10 * * * * *'
    cronDaily   : '0 10 0 * * *'
    run         : no
    defaultRecepient : undefined
  emailSender     :
    run           : no
  guests          :
    # define this to limit the number of guset accounts
    # to be cleaned up per collection cycle.
    poolSize      : 1e4
    batchSize     : undefined
    cleanupCron   : '*/10 * * * * *'
  pidFile         : '/tmp/koding.server.pid'
  loggr           :
    push          : no
    url           : ""
    apiKey        : ""
  librato         :
    push          : no
    email         : ""
    token         : ""
    interval      : 60000
  haproxy         :
    webPort       : 3020
  kontrold        :
    api           :
      port        : 8000
    proxy         :
      port        : 8080
      portssl     : 8081
      ftpip       : '127.0.0.1'
      sslips      : '127.0.0.1'
    rabbitmq      :
      host        : 'local.koding.com'
      port        : '5672'
      login       : 'guest'
      password    : 'guest'
      vhost       : '/'
  # crypto :
  #   encrypt: (str,key=Math.floor(Date.now()/1000/60))->
  #     crypto = require "crypto"
  #     str = str+""
  #     key = key+""
  #     cipher = crypto.createCipher('aes-256-cbc',""+key)
  #     cipher.update(str,'utf-8')
  #     a = cipher.final('hex')
  #     return a
  #   decrypt: (str,key=Math.floor(Date.now()/1000/60))->
  #     crypto = require "crypto"
  #     str = str+""
  #     key = key+""
  #     decipher = crypto.createDecipher('aes-256-cbc',""+key)
  #     decipher.update(str,'hex')
  #     b = decipher.final('utf-8')
  #     return b
  recurly       :
    apiKey      : 'b646d53c27e34916b7715931788df6af' # koding-test.recurly.com
  opsview       :
    push        : no
    host        : ''
  followFeed    :
    host        : 'local.koding.com'
    port        : 5672
    componentUser: 'guest'
    password    : 'guest'
    vhost       : 'followfeed'
