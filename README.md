Twilio Counter
=================

SMS:
---
- receiver.php
    - Add to the Twilio website (In SMS TwiML Area)
- server_maps.go
    - What you need to run the server, and point the receiver to. This will take the body content from receiver and increment it depending on how many times it's texted.
- server.go
    - The hard coded version of server_maps. This is if you don't need to be able to text anything - you input the values from before.
- twilio.go
    - Basic Twilio support demonstration


Works:
---
To run the server_maps.go server run: go run server_maps.go.

Setup your server IP in receiver.php. You have to alter "http://0.0.0.0:9999/" and set it to your server IP.

You have to go to 0.0.0.0:9999 or the port & IP you run the server on, and do /?sms=(NAME).

For example, my server runs on 0.0.0.0:9999, then to vote you must go to 0.0.0.0:9999/?sms=John, and it will increment it by one. 
