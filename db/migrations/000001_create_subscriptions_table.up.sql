CREATE TABLE IF NOT EXISTS subscriptions(
    subscription_id serial PRIMARY KEY,
    service_name VARCHAR (50) NOT NULL,
    price INTEGER NOT NULL,
    user_id UUID NOT NULL,
    start_date DATE NOT NULL,
    end_date DATE
);
