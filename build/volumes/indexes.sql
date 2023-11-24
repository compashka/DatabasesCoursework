CREATE INDEX idx_substation_location ON substations (location);

CREATE INDEX idx_request_worker_id ON requests (worker_id);

CREATE INDEX idx_transformer_substation ON transformers (substation);

CREATE INDEX idx_worker_name ON workers (first_name, last_name);

CREATE INDEX idx_transformer_factory_number ON transformers (factory_number);
