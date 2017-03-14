package main

// see the resource of : http://www.freecdn.cn/libs/highcharts/

// spline,line,column,area,bar
var TemplateSplineHtml = `{{define "T"}}
<!DOCTYPE HTML>
<html>
    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
        <title>Gochart - {{.ChartType}} | CodeG.cn</title>
<link rel="stylesheet" href="https://cdn.bootcss.com/bootstrap/3.3.7/css/bootstrap.min.css" integrity="sha384-BVYiiSIFeK1dGmJRAkycuHAHRg32OmUcww7on3RYdg4Va+PmSTsz/K68vbdEjh4u" crossorigin="anonymous">
        <script type="text/javascript" src="http://cdn.hcharts.cn/jquery/jquery-1.8.3.min.js"></script>
        <script type="text/javascript">
        $(function () {
            $('#container').highcharts({
                chart: {
                    // type: 'spline'
                    type: '{{.ChartType}}'
                                        // 默认是折线图(line)。type取值如下：
                                        // line:折线图
                                        // spline:曲线图，线条圆滑一些
                                        // column:竖向柱状图
                                        // area:面积图
                                        // bar:横向柱状图
                                        // 更多选项请参考：http://api.highcharts.com/highcharts#plotOptions
                },
                title: {
                    // text: 'Monthly Average Temperature',
                    text: '{{.Title}}',
                },
                subtitle: {
                    // text: 'Source: WorldClimate.com',
                    text: '{{.SubTitle}}',
                },
                xAxis: {
                    // categories: ['1', '2', '3', '4', '5', '6', '7', '8', '9', '10', '11', '12']
                    categories: [{{.XAxisNumbers}}] // JSONArray格式，数字，与x轴上点数的个数一致，
                },
                yAxis: {
                    title: {
                        // text: 'Temperature (°C)'
                        text: '{{.YAxisText}}'
                    },
                    plotLines: [{
                        value: 0,
                        width: 1,
                        color: '#808080'
                    }]
                },
                tooltip: {
                    shared: true,
                    // valueSuffix: '°C'
                    valueSuffix: '{{.ValueSuffix}}'
                },
                legend: {
                    layout: 'vertical',
                    align: 'right',
                    verticalAlign: 'middle',
                    borderWidth: 0
                },
                series: {{.DataArray}}
                /*
                [
	                // sample data is bellow
	                {
	                    name: 'Tokyo',
	                    data: [7.0, 6.9, 9.5, 14.5, 18.2, 21.5, 25.2, 26.5, 23.3, 18.3, 13.9, 9.6]
	                }, 
	                {
	                    name: 'New York',
	                    data: [-0.2, 0.8, 5.7, 11.3, 17.0, 22.0, 24.8, 24.1, 20.1, 14.1, 8.6, 2.5]
	                }, 
	                {
	                    name: 'Berlin',
	                    data: [-0.9, 0.6, 3.5, 8.4, 13.5, 17.0, 18.6, 17.9, 14.3, 9.0, 3.9, 1.0]
	                }, 
	                {
	                    name: 'London',
	                    data: [3.9, 4.2, 5.7, 8.5, 11.9, 15.2, 17.0, 16.6, 14.2, 10.3, 6.6, 4.8]
	                }
                ]
                */
            });
        });    
        </script>
    </head>
    <body>
    By <a id="copyright" class="anchor" href="http://blog.codeg.cn/2014/12/13/Hello-CodeG/" >zieckey@gmail.com</a>
    <script type="text/javascript" src="http://cdn.hcharts.cn/highcharts/4.0.1/highcharts.js"></script>
    <script type="text/javascript" src="http://cdn.hcharts.cn/highcharts/4.0.1/modules/exporting.js"></script>
    <script type="text/javascript" src="http://apps.bdimg.com/libs/bootstrap/3.3.4/js/bootstrap.min.js"></script>
    <div id="container" style="min-width: 310px; height: {{.Height}}px; margin: 0 auto"></div>

<script type ="text/javascript">
$(function(){
	var s = {{.DataArray}};
	var objstr = "";
	dataindex = 0;
	$(s).each(function(i,item){
		 dataindex = item.data.length;
		 var index = i+1;
		 objstr +='<tr><th scope="row">'+index+'</th><td>'+item.name+'</td>'
		 $(item.data).each(function(n,info){
		     objstr += '<td>'+info+'</td>';
		 });
		 objstr += "</tr>";
	});

	
	$("#tb").html(objstr);

})
</script>

<table id="tb"  class="table table-hover table-striped table-bordered">
  	
    </table>
    
</body>
</html>
{{end}}
`
