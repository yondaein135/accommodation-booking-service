INSERT INTO users (name, email, created_at, updated_at) VALUES
                                                            ('John Doe', 'john.doe@example.com', NOW(), NOW()),
                                                            ('Jane Smith', 'jane.smith@example.com', NOW(), NOW()),
                                                            ('Alice Johnson', 'alice.johnson@example.com', NOW(), NOW());

INSERT INTO billing_addresses (user_id, address_line1, address_line2, city, state, country, postal_code, created_at, updated_at) VALUES
                                                                                                                                     (1, '123 Main St', 'Apt 4', 'Springfield', 'State', 'Country', '12345', NOW(), NOW()),
                                                                                                                                     (2, '456 Elm St', 'Suite 5', 'Shelbyville', 'State', 'Country', '67890', NOW(), NOW()),
                                                                                                                                     (3, '789 Oak St', '', 'Centerville', 'State', 'Country', '54321', NOW(), NOW());



INSERT INTO properties (name, location, created_at, updated_at) VALUES
                                                                    ('Sea View Villa', 'Ocean Drive, Miami', NOW(), NOW()),
                                                                    ('Mountain Retreat', 'Highland Road, Denver', NOW(), NOW()),
                                                                    ('Urban Apartment', 'Downtown Street, New York', NOW(), NOW());



INSERT INTO rooms (property_id, type, capacity, price, created_at, updated_at) VALUES
                                                                                   (1, 'Deluxe Suite', 2, 300.00, NOW(), NOW()),
                                                                                   (1, 'Standard Room', 4, 180.00, NOW(), NOW()),
                                                                                   (2, 'Single Room', 1, 100.00, NOW(), NOW()),
                                                                                   (2, 'Double Room', 2, 150.00, NOW(), NOW()),
                                                                                   (3, 'Penthouse Suite', 3, 500.00, NOW(), NOW()),
                                                                                   (3, 'Studio Apartment', 2, 250.00, NOW(), NOW());

INSERT INTO bookings (user_id, room_id, start_date, end_date, price, occupants, is_paid, is_cancelled, created_at, updated_at) VALUES
                                                                                                                                   (1, 1, '2023-07-01', '2023-07-05', 500.00, 2, FALSE, FALSE, NOW(), NOW()),
                                                                                                                                   (2, 2, '2023-07-10', '2023-07-15', 750.00, 4, TRUE, FALSE, NOW(), NOW()),
                                                                                                                                   (3, 3, '2023-08-01', '2023-08-03', 300.00, 1, FALSE, TRUE, NOW(), NOW());


INSERT INTO payments (booking_id, amount, success, created_at, updated_at) VALUES
                                                                               ('1', 500.00, TRUE, NOW(), NOW()),
                                                                               ('2', 750.00, TRUE, NOW(), NOW()),
                                                                               ('3', 300.00, FALSE, NOW(), NOW());