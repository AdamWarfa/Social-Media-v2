-- Switch to the somev2_db database
\c somev2_db;

-- Create the users table
CREATE TABLE "users" (
  "id" VARCHAR(255) NOT NULL PRIMARY KEY,
  "avatar" TEXT,
  "followers" BIGINT,
  "password" VARCHAR(1024),
  "username" VARCHAR(1024),
  "email" VARCHAR(1024)
);

-- Insert data into the users table
INSERT INTO "users" (
  "id",
  "avatar",
  "followers",
  "password",
  "username",
  "email"
) VALUES
('-NYJioP0TPNf8sam-r_m','https://encrypted-tbn2.gstatic.com/images?q=tbn:ANd9GcTApJdGN6I8NjFvhKfOBIwjz759q53b8HBckMLXl_blfq5RXZFy',0,'123','Drake', 'drizzy@ovo.ca'),
('-NYK7iIF3ogIGvSuM_4y','https://upload.wikimedia.org/wikipedia/commons/thumb/5/5c/Kanye_West_at_the_2009_Tribeca_Film_Festival_%28crop_2%29.jpg/1200px-Kanye_West_at_the_2009_Tribeca_Film_Festival_%28crop_2%29.jpg',0,'test123','Kanye', 'yeezy@pablo.com'),
('-NYL-720SwuhqtJfN9a3','https://www.onthisday.com/images/people/michael-jordan.jpg?w=360',0,'123','MJ23', 'Jordan@nike.com'),
('98357e1d-b480-4a5a-bf2c-f55d280363ef','https://static.miraheze.org/loathsomecharacterswiki/thumb/7/7f/C--Users-cleme-Downloads-spongebob_PNG1.png/640px-C--Users-cleme-Downloads-spongebob_PNG1.png',0,'squidward','spongebob', 'spongie@bikini.com'),
('9920032e-23be-467f-b618-5ec00f5365a0','https://i.redd.it/3x8xbgq9gc351.jpg',0,'ibopro','Sheikh', 'haxibrim@nobody.com'),
('617ba003-ec2f-45bf-ae76-959f0d995fe5','https://i.guim.co.uk/img/media/3badb329c6e1ebe88ed95102daa034a83c7c1673/0_11_4000_2400/master/4000.jpg?width=1200&height=1200&quality=85&auto=format&fit=crop&s=56725b8827ee0e74ad7f3eb41d8575e9',0,'mahad123','Mahadinho', 'mahad@csharp.com');

-- Create the posts table
CREATE TABLE "posts" (
  "id" VARCHAR(120) NOT NULL PRIMARY KEY,
  "author" VARCHAR(120),
  "imgSrc" TEXT,
  "likes" BIGINT,
  "postDate" VARCHAR(200),
  "text" TEXT,
  FOREIGN KEY ("author") REFERENCES "users"("id")
);

-- Insert data into the posts table
INSERT INTO "posts" (
  "id",
  "author",
  "imgSrc",
  "likes",
  "postDate",
  "text"
) VALUES
('-NYJrGiacxi51YU0wTuD','-NYJioP0TPNf8sam-r_m','https://cdn.vox-cdn.com/thumbor/E-WhIDSz73yOwp0hes1YYtSk4Yc=/1400x1400/filters:format(jpeg)/cdn.vox-cdn.com/uploads/chorus_asset/file/15982917/usa_today_12406932.jpg',3,'2023-06-19T18:14:00.705Z','Jimmy buckets is him!!'),
('-NYJrM8VVxjM1kO18Vpm','-NYJioP0TPNf8sam-r_m','https://asset.dr.dk/imagescaler01/https%3A%2F%2Fwww.dr.dk%2Fimages%2Fother%2F2022%2F12%2F30%2Fscanpix-20221206-220030-3.jpg&w=1200&675&scaleAfter=crop',2,'2023-06-19T18:14:22.905Z','the goat'),
('-NYJrO4xSc42Ivtp3V1r','-NYJioP0TPNf8sam-r_m','https://upload.wikimedia.org/wikipedia/en/thumb/5/56/Real_Madrid_CF.svg/1200px-Real_Madrid_CF.svg.png',5,'2023-06-19T18:14:30.873Z','HALA MADRID'),
('-NYJyzAs8-20l-2sODdx','-NYJioP0TPNf8sam-r_m','https://upload.wikimedia.org/wikipedia/en/thumb/2/2d/Leicester_City_crest.svg/1200px-Leicester_City_crest.svg.png',1,'2023-06-19T18:47:41.938Z','rip leicester'),
('-NYK8dinGI3jlIgHCw3z','-NYK7iIF3ogIGvSuM_4y','https://cdn.britannica.com/39/7139-050-A88818BB/Himalayan-chocolate-point.jpg',1,'2023-06-19T19:34:17.686Z','My nigga since day 1'),
('-NYKzMvG3D7S-EkTwIi5','-NYK7iIF3ogIGvSuM_4y','https://pbs.twimg.com/media/FzAvYgPaYAArG7r?format=jpg&name=large',4,'2023-06-19T23:29:00.740Z','Unc Shay Shay'),
('-NYL-cJPKRU24XeFJhur','-NYL-720SwuhqtJfN9a3','https://upload.wikimedia.org/wikipedia/commons/b/b2/Hausziege_04.jpg',17,'2023-06-19T23:34:30.034Z','Me'),
('b6550227-b1cf-49a6-adcd-8d85ebede2f1','98357e1d-b480-4a5a-bf2c-f55d280363ef','https://upload.wikimedia.org/wikipedia/en/thumb/3/3b/SpongeBob_SquarePants_main_characters.png/370px-SpongeBob_SquarePants_main_characters.png',0,'2023-09-01T15:45:08.396Z','whole gang movin brazy'),
('ac5e25c5-73d7-4963-8208-66ffeff08b4a','98357e1d-b480-4a5a-bf2c-f55d280363ef','https://images.fineartamerica.com/images/artworkimages/mediumlarge/3/handsome-squidward-theodore-mitchell.jpg',0,'2023-09-01T16:16:13.378Z','bro wildin 💀💀');