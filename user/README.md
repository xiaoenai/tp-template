# user

Command user is the tp-micro service project.
<br>The framework reference: https://github.com/xiaoenai/tp-micro

## API Desc

### V1_User_Add

Add handler

- URI: `/user/v1/user/add`
- REQ-QUERY:
- REQ-BODY:

	```js
	{
		"name": ""	// {string} 
	}
	```

- RESULT:


### V1_User_GetById

获取用户

- URI: `/user/v1/user/get_by_id`
- REQ-QUERY:
	- `id={int64}`
- REQ-BODY:

	```js
	{}
	```

- RESULT:

	```js
	{
		"name": ""	// {string} 
	}
	```



### V1_User_Get

Get handler

- URI: `/user/v1/user/get`
- REQ-QUERY:
- REQ-BODY:
- RESULT:

	```js
	{
		"name": ""	// {string} 
	}
	```





<br>

*This is a project created by `micro gen` command.*

*[About Micro Command](https://github.com/xiaoenai/tp-micro/tree/v2/cmd/micro)*

## Error List

|Code|Message(输出时 Msg 将会被转为 JSON string)|
|------|------|
