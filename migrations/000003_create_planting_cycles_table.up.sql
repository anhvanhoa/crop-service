-- Create planting_cycles table
CREATE TABLE planting_cycles (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    cycle_name VARCHAR(255) NOT NULL,
    growing_zone_id uuid NOT NULL,
    plant_variety_id uuid NOT NULL,
    seed_date DATE,
    transplant_date DATE,
    expected_harvest_date DATE,
    actual_harvest_date DATE,
    plant_quantity INTEGER,
    seed_batch VARCHAR(100),
    status VARCHAR(50) DEFAULT 'planning', -- planning, seeding, transplanting, growing, flowering, harvesting, completed, failed
    notes TEXT,
    created_by uuid,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    FOREIGN KEY (plant_variety_id) REFERENCES plant_varieties(id)
);

-- Create indexes
CREATE INDEX idx_planting_cycles_zone_status ON planting_cycles (growing_zone_id, status);
CREATE INDEX idx_planting_cycles_variety ON planting_cycles (plant_variety_id);
CREATE INDEX idx_planting_cycles_dates ON planting_cycles (seed_date, expected_harvest_date);
CREATE INDEX idx_planting_cycles_created_by ON planting_cycles (created_by);

-- Create trigger for updated_at
CREATE TRIGGER update_planting_cycles_updated_at 
    BEFORE UPDATE ON planting_cycles 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column();
