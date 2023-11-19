Assumptions:
<li>
The application is running on a 64bit system
</li>

<h1>Notes</h1>
<li>
Run the application using ./run.sh. Once the db finished booting, the microservices will take care of defining the models.
You can track their progress on the console.
</li>
<li>
Once the db is set up and models are in place, insert some dummy data in by running the dummy.sql file in the root directory
</li>
<li>
You can print the booking history to the console by calling the GetBookingHistory handler on port 50052 (booking service), as per the requirements.
</li>


<h1>Test requests</h1>


1, List available rooms:

- ::50053/ListProperties 
- message:`{"desired_end_date":{"seconds":1699058400},"desired_start_date":{"seconds":1698947200},"price_max":900,"price_min":0}`

2, Create a booking

- ::50052/CreateBooking 
- message:`{"room_id":"1","start_date":{"seconds":"1732019439"},"end_date":{"seconds":"1732105839"},"occupants":2,"price":16,"user_id":"1"}`

3, Cancel your booking

- ::50052/CancelBooking
- message:`{"booking_id":"4","user_id":"1"}`

4, Get user's booking history

- ::50052/GetBookingHistory
- message:`{"price_max":10000000,"price_min":0,"start_date":{"seconds":1688169600},"user_id":"1"}`

5, Pay for a booking
- ::50054/ProcessPayment
- message:`{"amount":500.00,"booking_id":"1","payment_details":{"card_expiry":"12/23","card_number":"4111-1111-1111-1111","cvv":123,"first_name":"John","last_name":"Doe"}}`