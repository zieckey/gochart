package main

// see the resource of : http://www.freecdn.cn/libs/highcharts/

var TemplatePieHtml = `{{define "T"}}
<!DOCTYPE HTML>
<html>
    <head>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
        <title>Gochart - {{.ChartType}} | CodeG.cn</title>

        <script type="text/javascript" src="http://cdn.hcharts.cn/jquery/jquery-1.8.3.min.js"></script>
        <script type="text/javascript">
        $(function () {
            $('#container').highcharts({
                chart: {
                    //type: 'pie',
                    type: '{{.ChartType}}',
                    plotBackgroundColor: null,
                    plotBorderWidth: null,
                    plotShadow: false
                },
                title: {
                    // text: 'Browser market shares at a specific website, 2014'
                    text: '{{.Title}}',
                },
                subtitle: {
                    // text: 'Source: somewebsite.com',
                    text: '{{.SubTitle}}',
                },
                tooltip: {
                    pointFormat: '{series.name}: <b>{point.percentage:.1f}%</b>'
                },
                plotOptions: {
                    pie: {
                        allowPointSelect: true,
                        cursor: 'pointer',
                        dataLabels: {
                            enabled: true,
                            format: '<b>{point.name}</b>: {point.percentage:.1f} %',
                            style: {
                                color: (Highcharts.theme && Highcharts.theme.contrastTextColor) || 'black'
                            }
                        }
                    }
                },
                series: [{
                    // name: 'Browser share',
                    name : '{{.SeriesName}}',
                    data: 
                    	{{.DataArray}}
                    /*
                    data: 
                    [
                        ['Firefox',   45.0],
                        ['IE',       26.8],
                        ['Chrome',  12.8],
                        ['Safari',    8.5],
                        ['Opera',     6.2],
                        ['Others',   0.7]
                    ]
                    */
                }]
            });
        });
		</script>
    </head>
    <body>
    By <a id="copyright" class="anchor" href="http://blog.codeg.cn/2014/12/13/Hello-CodeG/" >zieckey@gmail.com</a>
    
    <script type="text/javascript" src="http://cdn.hcharts.cn/highcharts/4.0.1/highcharts.js"></script>
    <script type="text/javascript" src="http://cdn.hcharts.cn/highcharts/4.0.1/modules/exporting.js"></script>

    <div id="container" style="min-width: 310px; height: {{.Height}}px; max-width: 600px; margin: 0 auto"></div>

    </body>
</html>
{{end}}
`
