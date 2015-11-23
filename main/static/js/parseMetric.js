function parseMetric(influx_data) {
    var result = [];
    var series = influx_data[0].Series;
    var columns = influx_data[0].Series[0].columns;
    var legend = [];
    for (var i = 0; i < series.length; i++) {
        var name = series[i].name;
        legend.push(name);
        var values = series[i].values;
        var res = [];
        for (var j = 0; j < values.length; j++) {
            var obj1 = {};
            for (var k = 0; k < columns.length; k++) {
                if (columns[k] == "time") {
                    var d = new Date(values[j][k]);
                    values[j][k] = new Date(d.setTime( d.getTime() + 8*60*1000));
                }
                if (k == columns.length - 1) {
                    if (values[j][k] == null) {
                        values[j][k] = 0;
                    } else {
                        values[j][k] = parseFloat(values[j][k]);
                    }
                }
                obj1[columns[k]] = values[j][k];
            }
            res.push(obj1);
        }
        result.push(res);
    }
    return {'legend': legend, 'result': result};
}
