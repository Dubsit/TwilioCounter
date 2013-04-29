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
