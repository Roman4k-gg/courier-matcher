CREATE TABLE IF NOT EXISTS couriers (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    lat DOUBLE PRECISION NOT NULL,
    lon DOUBLE PRECISION NOT NULL,
    vehicle VARCHAR(20) NOT NULL,
    rating DOUBLE PRECISION DEFAULT 5.0,
    status VARCHAR(20) DEFAULT 'free',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO couriers (lat, lon, vehicle, rating, status)
SELECT 
    56.48 + (random() - 0.5) * 0.1,
    84.95 + (random() - 0.5) * 0.1,
    CASE (random() * 3)::int 
        WHEN 0 THEN 'car'
        WHEN 1 THEN 'bike'
        ELSE 'foot'
    END,
    3.5 + random() * 1.5,
    CASE WHEN random() < 0.7 THEN 'free' ELSE 'busy' END
FROM generate_series(1, 100)
ON CONFLICT DO NOTHING;
