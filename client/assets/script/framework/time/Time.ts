

export class Time {

    //将时间转成相应格式  by ming.lei 2018.06.06
    public static Format(time: Date, format = 'yyyy/MM/dd hh:mm:ss') {
        var o = {
            "M+": time.getMonth() + 1,                      //month 
            "d+": time.getDate(),                           //day 
            "h+": time.getHours(),                          //hour 
            "m+": time.getMinutes(),                        //minute 
            "s+": time.getSeconds(),                        //second 
            "q+": Math.floor((time.getMonth() + 3) / 3),    //quarter 
            "S": time.getMilliseconds()                     //millisecond 
        }

        if (/(y+)/i.test(format)) {
            format = format.replace(RegExp.$1, (time.getFullYear() + "").substr(4 - RegExp.$1.length));
        }

        for (var k in o) {
            if (new RegExp("(" + k + ")").test(format)) {
                format = format.replace(RegExp.$1, RegExp.$1.length == 1 ? o[k] : ("00" + o[k]).substr(("" + o[k]).length));
            }
        }
        return format;
    }

    //时间戳转时间
    public static TimestampToTime(timestamp: number) {
        return this.Format(new Date(timestamp))
    }

    
}