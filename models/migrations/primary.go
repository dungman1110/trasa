package migrations

// PrimaryMigration is main migration models. Always called in fresh installation of TRASA
// TODO init function to check and update database version. If version stored in database is below current version,
// find file name in migration package and execute database migration sql codes for that version. filename should be same as db version required.
var PrimaryMigration = []string{
	`CREATE TABLE IF NOT EXISTS db_version (
		version varchar PRIMARY KEY NOT NULL,
		created_at BIGINT
	);`,
	`CREATE TABLE IF NOT EXISTS org (
			id varchar PRIMARY KEY NOT NULL,
			org_name varchar NOT NULL,
			domain varchar,
			primary_contact varchar,
			timezone varchar,
			phone_number varchar,
			license JSONB,
			created_at BIGINT,
			CONSTRAINT unique_domain UNIQUE( domain),
			CONSTRAINT unique_org_name UNIQUE(org_name)
		);`,
	`CREATE TABLE IF NOT EXISTS users (
			org_id varchar NOT NULL,
			id varchar PRIMARY KEY NOT NULL,
			username varchar NOT NULL,
			first_name varchar,
			middle_name varchar,
			last_name varchar,
			email varchar NOT NULL,
			password varchar,
			user_role varchar,
			status BOOL,
			idp_name VARCHAR,
			external_id VARCHAR,
			public_key varchar,
			created_at BIGINT,
			updated_at BIGINT,
			deleted_at BIGINT,
			CONSTRAINT unique_email UNIQUE(org_id, email),
			CONSTRAINT unique_username UNIQUE(org_id, username)

		);`,
	`CREATE TABLE IF NOT EXISTS services (
			id varchar PRIMARY KEY NOT NULL,
			org_id varchar NOT NULL,
			name varchar,
			secret_key varchar,
			passthru BOOL,
			hostname varchar,
			type varchar,
			managed_accounts varchar,
			adhoc BOOL,
			remoteapp_name varchar,
			session_record BOOL,
			rdp_protocol varchar,
			native_log   BOOL,
			proxy_config JSONB,
			external_provider_name VARCHAR,
			external_id VARCHAR,
			external_security_group JSONB,
			distro_name VARCHAR,
			distro_version VARCHAR,
			ip_details JSONB,
			created_at BIGINT,
			updated_at BIGINT,
			deleted_at BIGINT,
			CONSTRAINT unique_servicename UNIQUE(org_id, name),
			CONSTRAINT unique_hostname UNIQUE(org_id, hostname, type)
		); `,

	`CREATE TABLE IF NOT EXISTS devices (
			id varchar PRIMARY KEY NOT NULL,
			user_id varchar NOT NULL REFERENCES users (id) ON DELETE CASCADE,
			org_id varchar NOT NULL,
			type varchar,
			device_hygiene JSONB,
			machine_id VARCHAR,
			fcm_token varchar NOT NULL,
			totpsec varchar,
			trusted BOOL,
			public_key varchar,
			added_at BIGINT,
			deleted BOOL
		);`,
	`CREATE TABLE IF NOT EXISTS browsers (
			id varchar primary key  NOT NULL,
			org_id VARCHAR NULL REFERENCES org  (id) ON DELETE CASCADE ,
			device_id varchar REFERENCES devices (id) ON DELETE CASCADE,
			user_agent varchar, 
			name varchar,
			version varchar,
			build varchar,
			extensions JSONB
		);`,
	`CREATE TABLE IF NOT EXISTS browser_ext (
			browser_id VARCHAR NOT NULL REFERENCES browsers (id) ON DELETE CASCADE,
			user_id VARCHAR NULL REFERENCES users (id) ON DELETE CASCADE ,
			org_id VARCHAR NULL REFERENCES org  (id) ON DELETE CASCADE ,
			ext_id VARCHAR NULL,
			name VARCHAR NULL,
			description VARCHAR NULL DEFAULT '':::STRING,
			version VARCHAR NULL,
			maydisable VARCHAR NULL,
			enabled BOOL NULL,
			install_type VARCHAR NULL,
			type VARCHAR NULL,
			perms VARCHAR[] NULL DEFAULT ARRAY[]:::VARCHAR[],
			host_perms VARCHAR[] NULL DEFAULT ARRAY[]:::VARCHAR[],
			isvuln BOOL NULL,
			vuln_reason VARCHAR NULL DEFAULT '':::STRING,
			last_checked INT8 NULL,
			UNIQUE INDEX unique_userextension (browser_id ASC, org_id ASC, user_id ASC, ext_id ASC)
		)`,
	`CREATE TABLE IF NOT EXISTS groups (
			id varchar PRIMARY KEY NOT NULL,
			org_id varchar REFERENCES org (id) ON DELETE CASCADE,
			name varchar ,
			type varchar,
			status BOOL,
			created_at BIGINT,
			updated_at BIGINT,
			CONSTRAINT unique_groupname UNIQUE(org_id,name)
		);`,
	`CREATE TABLE IF NOT EXISTS user_group_maps (
			id varchar PRIMARY KEY NOT NULL, 
			org_id varchar REFERENCES org(id) ON DELETE CASCADE,
			group_id varchar REFERENCES groups(id) ON DELETE CASCADE,
			user_id varchar REFERENCES users (id) ON DELETE CASCADE,
			status BOOL,
			created_at BIGINT,
			updated_at BIGINT,
			CONSTRAINT unique_useringroup UNIQUE(org_id,group_id,user_id)
		);`,
	`CREATE TABLE IF NOT EXISTS service_group_maps (
			id varchar PRIMARY KEY NOT NULL, 
			org_id varchar REFERENCES org (id) ON DELETE CASCADE,
			group_id varchar REFERENCES groups (id) ON DELETE CASCADE,
			service_id varchar REFERENCES services (id) ON DELETE CASCADE,
			status BOOL,
			created_at BIGINT,
			updated_at BIGINT,
			CONSTRAINT service_group_maps UNIQUE(org_id,group_id,service_id)
		);`,
	`CREATE TABLE IF NOT EXISTS policies (
			id varchar PRIMARY KEY NOT NULL,
			name varchar, 
			org_id varchar REFERENCES org(id) ON DELETE CASCADE,
			day_time JSONB NOT NULL,
			device_policy JSONB ,
			record_session BOOL,
			file_transfer BOOL,
			ip_source VARCHAR,
			risk_threshold FLOAT,
			tfa_enabled BOOL,
			expiry VARCHAR,
			allowed_countries VARCHAR,
			created_at BIGINT,
			updated_at BIGINT,
			CONSTRAINT unique_policyname UNIQUE(org_id,name)
		);`,

	`CREATE TABLE IF NOT EXISTS user_accessmaps (
			id varchar PRIMARY KEY NOT NULL,
			service_id varchar NOT NULL REFERENCES services (id) ON DELETE CASCADE,
			user_id varchar NOT NULL REFERENCES users (id) ON DELETE CASCADE,
			org_id varchar NOT NULL,
			policy_id varchar REFERENCES policies (id) ON DELETE CASCADE,
			privilege varchar,
			added_at BIGINT,
			CONSTRAINT unique_user_access_maps  UNIQUE(org_id,service_id, user_id, privilege)
		);`,

	`CREATE TABLE IF NOT EXISTS usergroup_accessmaps (
			id varchar PRIMARY KEY NOT NULL, 
			org_id varchar REFERENCES org (id) ON DELETE CASCADE,
			map_type varchar,
			servicegroup_id varchar ,
			usergroup_id varchar REFERENCES groups (id) ON DELETE CASCADE,
			policy_id varchar REFERENCES policies (id) ,
			privilege varchar,
			created_at BIGINT,
			CONSTRAINT unique_usergroup_accessmaps UNIQUE(org_id,servicegroup_id,usergroup_id, privilege)
		);`,
	`CREATE TABLE IF NOT EXISTS inapp_notifs (
			id varchar PRIMARY KEY NOT NULL,
			user_id varchar NOT NULL REFERENCES users (id) ON DELETE CASCADE,
			org_id varchar NOT NULL,
			emitter_id varchar NOT NULL  ,
			label varchar,
			text varchar,
			created_on BIGINT,
			is_resolved BOOL,
			resolved_on BIGINT
		);`,
	`CREATE TABLE IF NOT EXISTS gateway_http (
			id varchar PRIMARY KEY NOT NULL,
			org_id varchar NOT NULL,
			proxy_meta JSONB NOT NULL,
			service_id VARCHAR REFERENCES services (id) ON DELETE CASCADE,
			status BOOL,
			created_at BIGINT,
			updated_at BIGINT
		);`,
	`create table IF NOT EXISTS global_settings (
			id varchar PRIMARY KEY NOT NULL,
			org_id varchar REFERENCES org (id) ON DELETE CASCADE,
			status BOOL,
			type varchar,
			value JSONB,
			updated_by varchar,
			updated_on BIGINT,
			CONSTRAINT unique_globalsettings UNIQUE(org_id,type)
			)`,
	`create table IF NOT EXISTS policy_enforcer (
			id varchar PRIMARY KEY NOT NULL,
			user_id varchar REFERENCES users (id) ON DELETE CASCADE,
			org_id varchar REFERENCES org (id) ON DELETE CASCADE,
			type varchar,
			pending BOOL,
			assigned_by varchar,
			assigned_on BIGINT,
			resolved_on BIGINT,
			CONSTRAINT unique_enforcedpolicy UNIQUE(org_id,user_id,type)
			)`,
	`create table IF NOT EXISTS security_rules (
			id varchar PRIMARY KEY NOT NULL,
			org_id varchar REFERENCES org (id) ON DELETE CASCADE,
			name varchar,
			const_name varchar,
			description varchar,
			scope varchar,
			condition varchar,
			status BOOL,
			source varchar,
			action JSONB,
			created_by varchar,
			created_at BIGINT,
			last_modified BIGINT,
			CONSTRAINT unique_secrule UNIQUE(org_id,const_name)
			)`,
	`create table IF NOT EXISTS password_state (
			user_id varchar REFERENCES users (id) ON DELETE CASCADE,
			org_id varchar REFERENCES org (id) ON DELETE CASCADE,
			last_passwords varchar[],
			last_updated BIGINT,
			CONSTRAINT unique_passstate UNIQUE(user_id, org_id)
			)`,
	`create table if not exists auth_logs (
			id varchar PRIMARY KEY NOT NULL ,
			session_id varchar,
			access_device_id varchar,
			tfa_device_id varchar,
			service_id  varchar,
			service_name varchar,
			service_type varchar,
			email varchar,
			failed_reason varchar,
			geo_location jsonb,
			login_time BIGINT,
			logout_time BIGINT,
			org_id varchar ,
			server_ip varchar,
			session_duration int,
			status boolean,
			user_agent varchar,
			user_id varchar,
			user_ip varchar,
			privilege varchar,
			guests varchar[],
			recorded_session boolean
		)`,
	`create table if not exists inapp_trails (
			id varchar NOT NULL PRIMARY KEY,
			client_ip varchar,
			user_agent varchar,
			email varchar,
			event_time BIGINT,
			org_id varchar REFERENCES org (id) ON DELETE CASCADE ,
			description varchar,
			status boolean,
			user_id varchar
		);`,
	`create table if not exists signup_logs (
			company varchar,
			country varchar,
			email varchar,
			first_name varchar,
			last_name varchar,
			phone_number varchar,
			reference varchar,
			signup_time BIGINT
		);`,
	`create table IF NOT EXISTS service_keyvault (
			id VARCHAR PRIMARY KEY NOT NULL,
			org_id VARCHAR REFERENCES org (id) ON DELETE CASCADE ,
			service_id VARCHAR REFERENCES services(id) ON DELETE CASCADE ,
			secret_type VARCHAR,
			secret_id VARCHAR,
			secret BYTEA,
			added_at BIGINT,
			last_updated BIGINT,
			updated_by VARCHAR
		)`,
	`create table IF NOT EXISTS cert_holder (
		id VARCHAR PRIMARY KEY NOT NULL,
		org_id VARCHAR REFERENCES org (id) ON DELETE CASCADE ,
		entity_id VARCHAR ,
		cert BYTEA,
		key BYTEA,
		csr BYTEA,
		cert_type VARCHAR,
		created_at BIGINT,
		last_updated BIGINT,
		CONSTRAINT unique_cert UNIQUE(org_id,entity_id, cert_type)
	)`,
	`create table IF NOT EXISTS idp (
		id VARCHAR PRIMARY KEY NOT NULL,
		org_id VARCHAR REFERENCES org (id) ON DELETE CASCADE ,
		name VARCHAR NOT NULL,
		type VARCHAR NOT NULL,
		meta VARCHAR,
		is_enabled BOOL,
		client_id VARCHAR,
		redirect_url VARCHAR,
		audience_uri VARCHAR,
		endpoint VARCHAR,
		integration_type VARCHAR,
		scim_endpoint VARCHAR,
		created_by VARCHAR,
		last_updated BIGINT,
		CONSTRAINT unique_idp UNIQUE(org_id,type, name)
	)`,
	`create table IF NOT EXISTS key_holder (
		id VARCHAR PRIMARY KEY NOT NULL,
		org_id VARCHAR REFERENCES org (id) ON DELETE CASCADE ,
		name VARCHAR NOT NULL,
		tag VARCHAR,
		value BYTEA,
		added_by VARCHAR,
		added_at BIGINT,
		CONSTRAINT unique_keys UNIQUE(org_id,name)
	)`,
	`create table IF NOT EXISTS cloudiaas_sync (
		id VARCHAR PRIMARY KEY NOT NULL,
		org_id VARCHAR REFERENCES org (id) ON DELETE CASCADE ,
		name VARCHAR NOT NULL,
		last_synced_by VARCHAR NOT NULL,
		last_synced_on BIGINT,
		CONSTRAINT unique_cloudiaassync UNIQUE(org_id,name)
	)`,
	`create table IF NOT EXISTS keylog (
		id VARCHAR PRIMARY KEY NOT NULL,
		org_id VARCHAR REFERENCES org  (id) ON DELETE CASCADE ,
		hash VARCHAR NOT NULL,
		generated_at INT,
		status BOOL,
		last_updated BIGINT
	)`,
	`CREATE TABLE IF NOT EXISTS adhoc_perms (
			id varchar PRIMARY KEY NOT NULL,
			requester_id varchar NOT NULL REFERENCES users (id) ON DELETE CASCADE,
			org_id varchar NOT NULL,
			service_id varchar NOT NULL REFERENCES services (id) ON DELETE CASCADE,
			requestee_id varchar NOT NULL REFERENCES users (id) ON DELETE CASCADE,
			request_text varchar NOT NULL,
			requested_on BIGINT NOT NULL,
			is_authorized	BOOL,
			authorized_on BIGINT,
			authorized_period BIGINT,
			authorized_policy JSONB,
			is_expired BOOL,
			session_id varchar[]
		);`,

	`CREATE TABLE IF NOT EXISTS backups (
			id varchar PRIMARY KEY NOT NULL,
			org_id varchar NOT NULL,
			name varchar,
			type varchar,
			created_at BIGINT
		);`,
}
