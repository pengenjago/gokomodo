INSERT INTO public.buyers (id,created_at,updated_at,deleted_at,email,"name","password",address) VALUES 
('736b11bb-99b8-49ee-b48d-6937f6574cc5','2022-03-31 19:24:41.000','2022-03-31 19:24:41.000',NULL,'firmanelpri@gmail.com','Maanz','$2a$08$kZ.IAsglYvotP6l1Taoe6OkMVAhLSg6Mig8ty3TjhomdSeAelk0I.','Cakung')
,('a274ff30-9c18-4ca9-b3aa-9f5f8afc9ba9','2022-03-31 19:24:41.000','2022-03-31 19:24:41.000',NULL,'budiman@gmail.com','budz','$2a$08$W7iEG3dgLoP6NTU0WzypaulfhvZSRKv08Tkt8a7gJIuF576E8za4S','Tangerang')
;

INSERT INTO public.sellers (id,created_at,updated_at,deleted_at,email,"name","password",address) VALUES 
('qkwoqw-dmada-2122md','2022-03-31 19:24:41.000','2022-03-31 19:24:41.000',NULL,'deb@gmail.com','Toko Debby','$2a$08$a.Us/.iiURQeTZiTrQOgY.UA0.ZLP3llWxMyjYftR0k1pO1ev2U.S','Pondok Kopi')
,('99ade76c-dc6f-49b2-8c27-f41bbaca3528','2022-03-31 19:24:41.000','2022-03-31 19:24:41.000',NULL,'clothing@gmail.com','Toko Murah Meriah','$2a$08$a.Us/.iiURQeTZiTrQOgY.UA0.ZLP3llWxMyjYftR0k1pO1ev2U.S','Bandung')
;

INSERT INTO public.products (id,created_at,updated_at,deleted_at,product_name,description,price,seller_id) VALUES 
('430f4009-4c59-4aea-97bc-ddfc43ec3d67','2022-04-02 14:31:37.389','2022-04-02 14:31:37.389',NULL,'Samsung Galaxy S20','Samsung Galaxy S20 Internal 256GB',13000000,'qkwoqw-dmada-2122md')
,('44a6c61a-e820-4eea-bfea-13adffcbdad8','2022-04-02 15:00:06.596','2022-04-02 15:00:06.596',NULL,'Samsung Galaxy S20+','Samsung Galaxy S20+ 256GB',15000000,'qkwoqw-dmada-2122md')
,('370b5ed4-a910-4ab5-8fa1-9d2a27f39490','2022-04-02 17:26:19.745','2022-04-02 17:26:19.745',NULL,'Fjallraven - Foldsack No. 1 Backpack, Fits 15 Laptops','Your perfect pack for everyday use and walks in the forest. Stash your laptop (up to 15 inches) in the padded sleeve, your everyday',1580140,'99ade76c-dc6f-49b2-8c27-f41bbaca3528')
,('f0f23be1-bba4-41cd-a258-c09f26b6a219','2022-04-02 17:27:19.861','2022-04-02 17:27:19.861',NULL,'Mens Casual Premium Slim Fit T-Shirts','Slim-fitting style, contrast raglan long sleeve, three-button henley placket, light weight & soft fabric for breathable and comfortable wearing. And Solid stitched shirts with round neck made for durability and a great fit for casual fashion wear and diehard baseball fans. The Henley style round neckline includes a three-button placket.',320483,'99ade76c-dc6f-49b2-8c27-f41bbaca3528')
,('d81e5c72-4254-42d4-92a5-dbf32e00afb9','2022-04-02 17:28:00.680','2022-04-02 17:28:00.680',NULL,'Mens Cotton Jacket','Great outerwear jackets for Spring/Autumn/Winter, suitable for many occasions, such as working, hiking, camping, mountain/rock climbing, cycling, traveling or other outdoors. Good gift choice for you or your family member. A warm hearted love to Father, husband or son in this thanksgiving or Christmas Day.',804657,'99ade76c-dc6f-49b2-8c27-f41bbaca3528')
,('14de0816-ca6b-480d-acd5-28defbbbb844','2022-04-02 14:31:37.000','2022-04-02 14:31:37.000',NULL,'Iphone X','IPhone X  Internal 256 GB',7300000,'qkwoqw-dmada-2122md')
,('400ecf9e-e3e2-4e69-904c-b1b494c6105a','2022-04-02 14:31:37.000','2022-04-02 14:31:37.000',NULL,'Iphone X','IPhone X  Internal 128 GB',5800000,'qkwoqw-dmada-2122md')
,('9cf029c6-547b-4f5d-93e0-07390dbed56f','2022-04-02 14:31:37.000','2022-04-02 14:31:37.000',NULL,'Iphone X','IPhone X Internal 64Gb',5000000,'qkwoqw-dmada-2122md')
,('75dd87a0-ffee-4706-83da-5b392d71bae4','2022-04-03 22:18:34.426','2022-04-03 22:18:34.426',NULL,'Erigo Beach','Kaos Erigo Beach',80000,'99ade76c-dc6f-49b2-8c27-f41bbaca3528')
;

INSERT INTO public.orders (id,created_at,updated_at,deleted_at,buyer_id,seller_id,product_id,quantity,total_price,status) VALUES 
('9c1b7cd4-456a-48aa-ba2d-180b949be6aa','2022-04-03 21:55:32.949','2022-04-03 21:55:32.949',NULL,'a274ff30-9c18-4ca9-b3aa-9f5f8afc9ba9','qkwoqw-dmada-2122md','400ecf9e-e3e2-4e69-904c-b1b494c6105a',1,5800000,'Pending')
,('258fc335-7749-4c17-a9c9-e6be948933c7','2022-04-03 22:20:11.264','2022-04-03 22:20:11.264',NULL,'a274ff30-9c18-4ca9-b3aa-9f5f8afc9ba9','qkwoqw-dmada-2122md','400ecf9e-e3e2-4e69-904c-b1b494c6105a',1,5800000,'Pending')
,('6ceea581-94f0-4130-8435-250756224355','2022-04-03 22:20:11.444','2022-04-03 22:20:11.444',NULL,'a274ff30-9c18-4ca9-b3aa-9f5f8afc9ba9','qkwoqw-dmada-2122md','44a6c61a-e820-4eea-bfea-13adffcbdad8',1,15000000,'Pending')
,('bbb757d7-1bb6-46b7-b4f8-625f4d22966d','2022-04-03 21:55:33.212','2022-04-03 22:05:08.162',NULL,'a274ff30-9c18-4ca9-b3aa-9f5f8afc9ba9','qkwoqw-dmada-2122md','44a6c61a-e820-4eea-bfea-13adffcbdad8',1,15000000,'Pending')
;