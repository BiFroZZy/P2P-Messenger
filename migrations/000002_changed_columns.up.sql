alter table messenger_users alter column user_login type varchar(20);
alter table messenger_users alter column user_password type varchar(100);

alter table messenger_users alter column user_login type varchar(60);
alter table messenger_users alter column drop user_email;
