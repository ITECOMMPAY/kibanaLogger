# kibanaLogger
Simple kibana logger.  
You can specify log file with flag
`-kibanaLog=filename`

###Usage
`kibanaLogger.Info("message")` write log message with status **OK**  
`kibanaLogger.Warning("message")` write log message with status **WRN**  
`kibanaLogger.Error("message")` write log message with status **ERR**
