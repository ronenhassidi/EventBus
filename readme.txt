
RabbitMQ
================
http://localhost:15672

There are 2 main 
1) Controller\main.go
2) Agent\main.go

(*) you need to run both the controller and agent. 
Controller
----------
    The controller creates all the queues that are defined in controller.yaml and produce messages.
    It also starts a service that can get API calls that sends the message to a queue
Agent
------- 
    The Agent consumes the messages for all the queues that are defined in agent.yaml 


sample
-----------
{
  "msgId": "1",
  "station": "station1",
  "msgType": "API_TYPE",
  "apiMsg": {
      "body":"aaa",
      "methodType":"bbb",
      "url":"cccc",
      "apiResponse":{
         "body":"ddd", 
         "headers": 
             [
                 {
             "key":"eee",
             "value":"fff"
                 },
                 {
             "key":"ggg",
             "value":"hhh"
                 }
             ]
      }
   },
   "dbMsg": {},
   "fsMsg": {}
}
