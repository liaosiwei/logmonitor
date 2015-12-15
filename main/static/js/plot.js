function plot(begin, end) {
    $.ajax({
    url: "/query/static/?db=webproxy&from=" + begin + "&to=" + end,
    type: "GET",
    contentType: "application/json; charset=utf-8",
    dataType: "json",
    async: true,
    success: function (data) {
        console.log(data);
        jsonData = parseMetric(data);

        MG.data_graphic({
            title: "Webproxy 90 Percentile",
            description: "webproxy tm, btm, ctm and size 90 percentile data plot",
            data: jsonData.result,
            area: false,
            width: 580,
            height: 300,
            target: '#webproxy',
            legend: jsonData.legend,
            legend_target: '#web_legend',
            x_accessor: 'time',
            y_accessor: 'value',
            baselines: [{value: 4, label: 'baseline'}],
        })    
    },
    error: function(result) {
        console.log("get tsv data failed!");
    }
    }).done(function() {console.log("done!");});

    $.ajax({
    url: "/query/static/?db=ac&from=" + begin + "&to=" + end,
    type: "GET",
    contentType: "application/json; charset=utf-8",
    dataType: "json",
    async: true,
    success: function (data) {
        console.log(data);
        jsonData = parseMetric(data);

        MG.data_graphic({
            title: "Ac 90 Percentile",
            description: "Ac tm, GT, FT, lbt 90 percentile point plot",
            data: jsonData.result,
            area: false,
            width: 580,
            height: 300,
            target: '#ac',
            legend: jsonData.legend,
            legend_target: '#ac_legend',
            x_accessor: 'time',
            y_accessor: 'value',
            baselines: [{value: 2, label: 'baseline'}],
        })    
    },
    error: function(result) {
        console.log("get tsv data failed!");
    }
    }).done(function() {console.log("done!");});

    $.ajax({
    url: "/query/static/?db=attr&from=" + begin + "&to=" + end,
    type: "GET",
    contentType: "application/json; charset=utf-8",
    dataType: "json",
    async: true,
    success: function (data) {
        console.log(data);
        jsonData = parseMetric(data);

        MG.data_graphic({
            title: "Attr 90 Percentile",
            description: "Attr tm, dt, fct 90 percentile point plot",
            data: jsonData.result,
            area: false,
            width: 580,
            height: 300,
            target: '#attr',
            legend: jsonData.legend,
            legend_target: '#attr_legend',
            x_accessor: 'time',
            y_accessor: 'value',
            baselines: [{value: 1, label: 'baseline'}],
        })    
    },
    error: function(result) {
        console.log("get tsv data failed!");
    }
    }).done(function() {console.log("done!");});

}
