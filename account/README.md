# account

Command account is the tp-micro service project.
<br>The framework reference: https://github.com/xiaoenai/tp-micro

## API Desc

### V1_User_Set

增加用户

- URI: `/account/v1/user/set`
- REQ-QUERY:
- REQ-BODY:

	```js
	{
		"name": ""	// {string} 
	}
	```

- RESULT:


### V1_User_GetById

根据ID获取user

- URI: `/account/v1/user/get_by_id`
- REQ-QUERY:
- REQ-BODY:

	```js
	{
		"id": -0	// {int64} 
	}
	```

- RESULT:

	```js
	{
		"id": -0,	// {int64} 
		"name": "",	// {string} 
		"updated_at": -0,	// {int64} 
		"created_at": -0,	// {int64} 
		"deleted_ts": -0	// {int64} 
	}
	```





<br>

*This is a project created by `micro gen` command.*

*[About Micro Command](https://github.com/xiaoenai/tp-micro/tree/v2/cmd/micro)*

## Error List

|Code|Message(输出时 Msg 将会被转为 JSON string)|
|------|------|
