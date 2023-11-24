INSERT INTO substations (name, location, year_of_construction, commissioning_year)
VALUES ('Пальники', 'Залесная', 2011, 2012), ('Гари', 'Гари', 2003, 2004);

INSERT INTO factories (name, city) VALUES ('ТТ3', 'Тольятти');

INSERT INTO range_of_high_voltage_equipments (high_voltage_switch, medium_voltage_switch, low_voltage_switch)
VALUES ('Высоковольтный_A', 'Средневольтный_A', 'Низковольтный_A'),
       ('Высоковольтный_ь_B', 'Средневольтный_B', 'Низковольтный_B');
INSERT INTO cable_lines (mark) VALUES ('КЛ_1'), ('КЛ_2');

INSERT INTO tire_sections (name) VALUES ('СШ_1'), ('СШ_2');

INSERT INTO cell_kvls (dispatch_name, cable_line, current_transformer, switch, protection_transformer, tire_section, number_of_current_transformers)
VALUES ('Ячейка_КВЛ_1', 'КЛ_1', 'Т_тока_A', 'Выключатель_A', 'Т_защиты_A', 'СШ_1', 3),
       ('Ячейка_КВЛ_2', 'КЛ_2', 'Т_тока_B', 'Выключатель_B', 'Т_защиты_B', 'СШ_2', 4);

INSERT INTO fuses (mark) VALUES ('ПН_1'), ('ПН_2');

INSERT INTO cell_tns (dispatch_name, voltage_transformer, fuse, tire_section)
VALUES ('Ячейка_ТН_1', 'Т_напряжения_A', 'ПН_1', 'СШ_1'),
       ('Ячейка_ТН_2', 'Т_напряжения_B', 'ПН_2', 'СШ_2');

INSERT INTO cell_tsns (dispatch_name, auxiliary_transformer, fuse, tire_section)
VALUES ('Ячейка_ТСН_1', 'Вспомогательный_A', 'ПН_1', 'СШ_1'),
       ('Ячейка_ТСН_2', 'Вспомогательный_B', 'ПН_2', 'СШ_2');

INSERT INTO nsses (rated_voltage_kV) VALUES (110), (220);

INSERT INTO range_of_standard_voltages (rated_winding_voltage_HV_kV, rated_winding_voltage_MV_kV, rated_winding_voltage_LV_kV)
VALUES (110, 10, 0.4), (220, 20, 0.4);

INSERT INTO type_of_transformers (type, power_MVA, cooling_system_type, range_of_standard_voltage)
VALUES ('Трансформатор_Т1', 100, 'Воздушное', 1), ('Трансформатор_Т2', 200, 'Водяное', 2);

INSERT INTO transformers (factory_number, NSS, substation, factory, type, date_of_manufacture, commissioning_date, dispatch_name, range_of_high_voltage_equipment, tire_section)
VALUES (1001, 1, 'Пальники', 'ТТ3', 'Трансформатор_Т1', '2020-02-15', '2020-03-01', 'Ячейка_ТН_1', 1, 'СШ_1'),
       (2002, 2, 'Гари', 'ТТ3', 'Трансформатор_Т2', '2018-05-20', '2018-06-05', 'Ячейка_ТН_2', 2, 'СШ_2');

INSERT INTO workers (first_name, last_name) VALUES ('Алексей', 'Иванов'), ('Екатерина', 'Смирнова');

INSERT INTO users (username, password, role) VALUES ('dispatcher1', 'dispatcher1', 'DISPATCHER'), ('worker1', 'worker1', 'WORKER');

