-- Create plant_varieties table
CREATE TABLE plant_varieties (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    scientific_name VARCHAR(255),
    category VARCHAR(100), -- leafy_vegetable, fruit_vegetable, herb, root_vegetable
    growing_season VARCHAR(100), -- spring, summer, autumn, winter, year_round
    growth_duration_days INTEGER, -- Số ngày từ gieo đến thu hoạch
    optimal_temp_min DECIMAL(5,2), -- Nhiệt độ tối thiểu (°C)
    optimal_temp_max DECIMAL(5,2), -- Nhiệt độ tối đa (°C)
    optimal_humidity_min DECIMAL(5,2), -- Độ ẩm tối thiểu (%)
    optimal_humidity_max DECIMAL(5,2), -- Độ ẩm tối đa (%)
    ph_min DECIMAL(3,1), -- pH tối thiểu
    ph_max DECIMAL(3,1), -- pH tối đa
    water_requirement VARCHAR(50), -- low, medium, high
    light_requirement VARCHAR(50), -- low, medium, high, full_sun
    description TEXT,
    media_id uuid,
    status VARCHAR(50) DEFAULT 'active',
    created_by uuid,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Create indexes
CREATE INDEX idx_plant_varieties_category ON plant_varieties (category);
CREATE INDEX idx_plant_varieties_season ON plant_varieties (growing_season);
CREATE INDEX idx_plant_varieties_status ON plant_varieties (status);
CREATE INDEX idx_plant_varieties_created_by ON plant_varieties (created_by);

-- Create trigger for updated_at
CREATE TRIGGER update_plant_varieties_updated_at 
    BEFORE UPDATE ON plant_varieties 
    FOR EACH ROW 
    EXECUTE FUNCTION update_updated_at_column();
