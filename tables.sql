CREATE TABLE banners (
  id INT PRIMARY KEY,
  feature_id INT NOT NULL,
  tag_id INT NOT NULL,
  included BOOLEAN DEFAULT false,
  UNIQUE (feature_id, tag_id)
);


