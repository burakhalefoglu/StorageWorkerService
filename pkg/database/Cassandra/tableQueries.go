package cassandra

var tableQueries = [21]string{
	`CREATE TABLE IF NOT EXISTS client_database.adv_events(id bigint, client_id bigint, project_id bigint, customer_id bigint, level_name text, level_index int, adv_type text, in_minutes decimal, triggered_time timestamp, status boolean, PRIMARY KEY(id))`,
	`CREATE TABLE IF NOT EXISTS client_database.buying_events(id bigint, client_id bigint, project_id bigint, customer_id bigint, level_name text, level_index int, product_type text, in_minutes decimal, triggered_time timestamp, status boolean, PRIMARY KEY(id))`,
	`CREATE TABLE IF NOT EXISTS client_database.enemy_base_login_levels(id bigint, client_id bigint, project_id bigint, customer_id bigint, level_name text, level_index int, playing_time int, average_scores int, date_time timestamp, is_dead tinyint, total_power_usage int,  status boolean, PRIMARY KEY(id))`,
	`CREATE TABLE IF NOT EXISTS client_database.enemy_base_level_fails(id bigint, client_id bigint, project_id bigint, customer_id bigint, fail_time_after_level_starting int, level_name text, level_index int, fail_location_x decimal, fail_location_y decimal, fail_location_z decimal, date_time timestamp,  status boolean, PRIMARY KEY(id))`,
	`CREATE TABLE IF NOT EXISTS client_database.game_sessions(id bigint, client_id bigint, project_id bigint, customer_id bigint, session_start_time timestamp, session_finish_time timestamp, session_time decimal, created_at timestamp, status boolean, PRIMARY KEY ((project_id), created_at)) WITH CLUSTERING ORDER BY (created_at DESC);`,
	`CREATE TABLE IF NOT EXISTS client_database.level_base_sessions(id bigint, client_id bigint, project_id bigint, customer_id bigint, level_name text, level_index int, session_time_minute decimal, session_start_time timestamp, session_finish_time timestamp, status boolean, PRIMARY KEY(id))`,
	`CREATE TABLE IF NOT EXISTS client_database.clients(id bigint, project_id bigint, is_paid_client boolean, created_at timestamp, paid_time timestamp, status boolean, PRIMARY KEY ((project_id),created_at)) WITH CLUSTERING ORDER BY (created_at DESC);`,
	`CREATE TABLE IF NOT EXISTS client_database.churn_bloker_ml_results(id bigint, client_id bigint, project_id bigint, customer_id bigint, model_type text, model_result double, date_time timestamp, status boolean, PRIMARY KEY(id))`,
	`CREATE TABLE IF NOT EXISTS client_database.offer_behaviors(id bigint, client_id bigint, project_id bigint, customer_id bigint, version smallint, offer_id int, isBuy_offer tinyint, created_at timestamp, status boolean, PRIMARY KEY ((project_id, offer_id), created_at)) WITH CLUSTERING ORDER BY (created_at DESC);`,
	`CREATE TABLE IF NOT EXISTS client_database.churn_prediction_ml_results(id bigint, client_id bigint, project_id bigint, customer_id bigint, model_type text, model_result decimal, date_time timestamp, status boolean, PRIMARY KEY(id))`,
	`CREATE TABLE IF NOT EXISTS client_database.hardwares(id bigint, client_id bigint, project_id bigint, customer_id bigint, device text, device_name text, device_type int, graphics_device_name text, graphics_device_type int, graphics_device_vendor text, graphics_device_version text, graphics_memory_size int, operating_system text, processor_count int, processor_type text, system_memory_size int, date_time timestamp, status boolean, PRIMARY KEY(id))`,
	`CREATE TABLE IF NOT EXISTS client_database.manuel_flows(id bigint, client_id bigint, project_id bigint, customer_id bigint, difficulty_level int, date_time timestamp, status boolean, PRIMARY KEY(id))`,
	`CREATE TABLE IF NOT EXISTS client_database.screen_clicks(id bigint, client_id bigint, project_id bigint, customer_id bigint, start_loc_x decimal, start_loc_y decimal, finish_loc_x decimal, finish_loc_y decimal, level_name text, level_index int, tab_count int, finger_id int, created_at timestamp, status boolean, PRIMARY KEY(id))`,
	`CREATE TABLE IF NOT EXISTS client_database.screen_swipes(id bigint, client_id bigint, project_id bigint, customer_id bigint, start_loc_x decimal, start_loc_y decimal, finish_loc_x decimal, finish_loc_y decimal, level_name text, level_index int, swipe_direction int, date_time timestamp, status boolean, PRIMARY KEY(id))`,
	`CREATE TABLE IF NOT EXISTS client_database.locations(id bigint, client_id bigint, project_id bigint, customer_id bigint, continent text, country text, city text, query text, region text, org text, created_at timestamp, status boolean, PRIMARY KEY(id))`,
	`CREATE TABLE IF NOT EXISTS client_database.inventorys(id bigint, client_id bigint, project_id bigint, customer_id bigint, minor_mine decimal, moderate_mine decimal, precious_mine decimal, created_at timestamp, status boolean, PRIMARY KEY(id))`,
	`CREATE TABLE IF NOT EXISTS client_database.items(id bigint, inventory_id bigint, item_type text, count int, PRIMARY KEY(id))`,
	`CREATE TABLE IF NOT EXISTS client_database.skills(id bigint, inventory_id bigint, skill_type text, count int, PRIMARY KEY(id))`,
	`CREATE TABLE IF NOT EXISTS client_database.temporary_abilitys(id bigint, inventory_id bigint, ability_type text, count int, PRIMARY KEY(id))`,
	`CREATE TABLE IF NOT EXISTS client_database.churn_prediction_success_rates(id bigint, project_id bigint, value float, status boolean, created_at timestamp, PRIMARY KEY ((project_id),created_at)) WITH CLUSTERING ORDER BY (created_at DESC);`,
	`CREATE TABLE IF NOT EXISTS client_database.adv_strategies(id bigint, project_id bigint, client_id bigint, strategy_id bigint, name text, version int, status boolean, created_at timestamp, PRIMARY KEY ((project_id, strategy_id),created_at)) WITH CLUSTERING ORDER BY (created_at DESC);`,
}

func GetTableQueries() [21]string {
	return tableQueries
}
