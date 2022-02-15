package cassandra

var tableQueries = [19]string{
	`CREATE TABLE IF NOT EXISTS ClientDatabase.adv_events(id bigint, client_id bigint, project_id bigint, customer_id bigint, level_name text, level_index int, adv_type text, in_minutes decimal, triggered_time date, status boolean, PRIMARY KEY(id))`,
	`CREATE TABLE IF NOT EXISTS ClientDatabase.buying_events(id bigint, client_id bigint, project_id bigint, customer_id bigint, level_name text, level_index int, product_type text, in_minutes decimal, triggered_time date, status boolean, PRIMARY KEY(id))`,
	`CREATE TABLE IF NOT EXISTS ClientDatabase.enemy_base_login_level_models(id bigint, client_id bigint, project_id bigint, customer_id bigint, level_name text, level_index int, playing_time int, average_scores int, date_time date, is_dead tinyint, total_power_usage int,  status boolean, PRIMARY KEY(id))`,
	`CREATE TABLE IF NOT EXISTS ClientDatabase.enemy_base_level_fail_models(id bigint, client_id bigint, project_id bigint, customer_id bigint, fail_time_after_level_starting int, level_name text, level_index int, fail_location_x decimal, fail_location_y decimal, fail_location_z decimal, date_time date,  status boolean, PRIMARY KEY(id))`,
	`CREATE TABLE IF NOT EXISTS ClientDatabase.game_session_models(id bigint, client_id bigint, project_id bigint, customer_id bigint, session_start_time date, session_finish_time date, session_time_minute decimal, status boolean, PRIMARY KEY(id))`,
	`CREATE TABLE IF NOT EXISTS ClientDatabase.level_base_session_models(id bigint, client_id bigint, project_id bigint, customer_id bigint, level_name text, level_index int, session_time_minute decimal, session_start_time date, session_finish_time date, status boolean, PRIMARY KEY(id))`,
	`CREATE TABLE IF NOT EXISTS ClientDatabase.client_data_models(id bigint, project_id bigint, is_paid_client tinyint, created_at date, paid_time date,  status boolean, PRIMARY KEY(id))`,
	`CREATE TABLE IF NOT EXISTS ClientDatabase.churn_bloker_ml_results(id bigint, client_id bigint, project_id bigint, customer_id bigint, model_type text, model_result double, date_time date, status boolean, PRIMARY KEY(id))`,
	`CREATE TABLE IF NOT EXISTS ClientDatabase.offer_behavior_models(id bigint, client_id bigint, project_id bigint, customer_id bigint, version smallint, offer_id int, date_time date, isBuy_offer tinyint, status boolean, PRIMARY KEY(id))`,
	`CREATE TABLE IF NOT EXISTS ClientDatabase.churn_prediction_ml_result_models(id bigint, client_id bigint, project_id bigint, customer_id bigint, model_type text, model_result decimal, date_time date, status boolean, PRIMARY KEY(id))`,
	`CREATE TABLE IF NOT EXISTS ClientDatabase.hardware_models(id bigint, client_id bigint, project_id bigint, customer_id bigint, device_model text, device_name text, device_type int, graphics_device_name text, graphics_device_type int, graphics_device_vendor text, graphics_device_version text, graphics_memory_size int, operating_system string, processor_count int, processor_type text, system_memory_size int, date_time date, status boolean, PRIMARY KEY(id))`,
	`CREATE TABLE IF NOT EXISTS ClientDatabase.manuel_flow_models(id bigint, client_id bigint, project_id bigint, customer_id bigint, difficulty_level int, date_time date, status boolean, PRIMARY KEY(id))`,
	`CREATE TABLE IF NOT EXISTS ClientDatabase.screen_click_models(id bigint, client_id bigint, project_id bigint, customer_id bigint, start_loc_x decimal, start_loc_y decimal, finish_loc_x decimal, finish_loc_y decimal, level_name text, level_index int, tab_count int, finger_id int, created_at date, status boolean, PRIMARY KEY(id))`,
	`CREATE TABLE IF NOT EXISTS ClientDatabase.screen_swipe_models(id bigint, client_id bigint, project_id bigint, customer_id bigint, start_loc_x decimal, start_loc_y decimal, finish_loc_x decimal, finish_loc_y decimal, level_name text, level_index int, swipe_direction int, date_time date, status boolean, PRIMARY KEY(id))`,
	`CREATE TABLE IF NOT EXISTS ClientDatabase.location_models(id bigint, client_id bigint, project_id bigint, customer_id bigint, continent text, country text, city text, query text, region text, org text, created_at date, status boolean, PRIMARY KEY(id))`,
	`CREATE TABLE IF NOT EXISTS ClientDatabase.inventory_models(id bigint, client_id bigint, project_id bigint, customer_id bigint, minor_mine decimal, moderate_mine decimal, precious_mine decimal, created_at date, status boolean, PRIMARY KEY(id))`,
	`CREATE TABLE IF NOT EXISTS ClientDatabase.item_models(id bigint, inventory_id bigint, item_type text, count int, PRIMARY KEY(id))`,
	`CREATE TABLE IF NOT EXISTS ClientDatabase.skill_models(id bigint, inventory_id bigint, skill_type text, count int, PRIMARY KEY(id))`,
	`CREATE TABLE IF NOT EXISTS ClientDatabase.temporary_ability_models(id bigint, inventory_id bigint, ability_type text, count int, PRIMARY KEY(id))`,
}

func GetTableQueries() [19]string {
	return tableQueries
}
