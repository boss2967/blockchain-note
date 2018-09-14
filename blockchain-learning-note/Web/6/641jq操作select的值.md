1. 获取第一个option的值       

	```
	$('#test option:first').val();
	```

2. 最后一个option的值                     

	```
	$('#test option:last').val();
	```

3. 获取第二个option的值          

	```
	$('#test option:eq(1)').val();
	``` 

4. 获取选中的值                         

	```
	$('#test').val();
	
	$('#test option:selected').val();
	```

5. 设置值为2的option为选中状态   

	```
	$('#test').attr('value','2');
	```

6. 设置最后一个option为选中

	```
	$('#test option:last').attr('selected','selected');
	
	$("#test").attr('value' , $('#test option:last').val());
	
	$("#test").attr('value' , $('#test option').eq($('#test option').length - 1).val());
	```

7. 获取select的长度            
	
	```
	$('#test option').length;
	``` 

8. 添加一个option
	
	```
	$("#test").append("<option value='n+1'>第N+1项</option>");
	
	$("<option value='n+1'>第N+1项</option>").appendTo("#test");
	```

9. 添除选中项       

	```
	$('#test option:selected').remove();
	``` 

10. 删除项选中(这里删除第一项)       

	```
	$('#test option:first').remove();
	```

11. 指定值被删除

	```
	$('#test option').each(function(){
	
	   if( $(this).val() == '5'){
	
	        $(this).remove();
	    }
	});
	
	$('#test option[value=5]').remove();
	```

12. 获取第一个Group的标签

	```
	$('#test optgroup:eq(0)').attr('label');
	```

13. 获取第二group下面第一个option的值

	```
	$('#test optgroup:eq(1) : option:eq(0)').val();
	```

14. 根据option的值选中option

	```
	$("#sel option:contains('C')").prop("selected", true);
	```