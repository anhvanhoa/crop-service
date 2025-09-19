-- 3. BẢNG DANH MỤC GIỐNG RAU
CREATE TABLE plant_varieties (
    id VARCHAR(36) PRIMARY KEY DEFAULT (UUID()),
    name VARCHAR(255) NOT NULL,
    scientific_name VARCHAR(255),
    category VARCHAR(100) COMMENT 'leafy_vegetable, fruit_vegetable, herb, root_vegetable',
    growing_season VARCHAR(100) COMMENT 'spring, summer, autumn, winter, year_round',
    growth_duration_days INTEGER COMMENT 'Số ngày từ gieo đến thu hoạch',
    optimal_temp_min DECIMAL(5,2) COMMENT 'Nhiệt độ tối thiểu (°C)',
    optimal_temp_max DECIMAL(5,2) COMMENT 'Nhiệt độ tối đa (°C)',
    optimal_humidity_min DECIMAL(5,2) COMMENT 'Độ ẩm tối thiểu (%)',
    optimal_humidity_max DECIMAL(5,2) COMMENT 'Độ ẩm tối đa (%)',
    ph_min DECIMAL(3,1) COMMENT 'pH tối thiểu',
    ph_max DECIMAL(3,1) COMMENT 'pH tối đa',
    water_requirement VARCHAR(50) COMMENT 'low, medium, high',
    light_requirement VARCHAR(50) COMMENT 'low, medium, high, full_sun',
    description TEXT,
    image_url VARCHAR(36),
    status VARCHAR(50) DEFAULT 'active',
    created_by VARCHAR(36),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    FOREIGN KEY (created_by) REFERENCES users(id),
    FOREIGN KEY (image_url) REFERENCES media(id),
    INDEX idx_plant_varieties_category (category),
    INDEX idx_plant_varieties_season (growing_season),
    INDEX idx_plant_varieties_status (status)
);

-- 4. BẢNG CHU KỲ TRỒNG
CREATE TABLE planting_cycles (
    id VARCHAR(36) PRIMARY KEY DEFAULT (UUID()),
    cycle_name VARCHAR(255) NOT NULL,
    growing_zone_id VARCHAR(36) NOT NULL,
    plant_variety_id VARCHAR(36) NOT NULL,
    seed_date DATE,
    transplant_date DATE,
    expected_harvest_date DATE,
    actual_harvest_date DATE,
    plant_quantity INTEGER,
    seed_batch VARCHAR(100),
    status VARCHAR(50) DEFAULT 'planning' COMMENT 'planning, seeding, transplanting, growing, flowering, harvesting, completed, failed',
    notes TEXT,
    created_by VARCHAR(36),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    FOREIGN KEY (growing_zone_id) REFERENCES growing_zones(id) ON DELETE CASCADE,
    FOREIGN KEY (plant_variety_id) REFERENCES plant_varieties(id),
    FOREIGN KEY (created_by) REFERENCES users(id),
    INDEX idx_planting_cycles_zone_status (growing_zone_id, status),
    INDEX idx_planting_cycles_variety (plant_variety_id),
    INDEX idx_planting_cycles_dates (seed_date, expected_harvest_date)
);