import $ from 'jquery';
import VueCookies from "vue-cookies";
var __enums = {};

export function EnumUtility() {}

EnumUtility.prototype.Set = function (data) {
    if (!Array.isArray(data)) {
        console.error("枚举输入参数必须是数组");
        return false;
    }
    if (data.length == 0) {
        console.error("枚举输入参数不能为空");
        return false;
    }
    for (var item in __enums) {
        console.error("枚举数据已存在,不允许重复赋值");
        return false;
    }
    if (sessionStorage.getItem("menus")) {
        return
    }
    data.forEach(function (item) {
        if (!__enums[item.type]) {
            __enums[item.type] = [];
        }
        __enums[item.type].push(item);
    });
    for(var key in __enums){
        sessionStorage.setItem("menus_"+key, JSON.stringify(__enums[key]));
    }
    return true;
};

EnumUtility.prototype.Get = function (type) {
    if (!type) return [];
    if(!sessionStorage.getItem("menus_"+type)){
        console.log("没有枚举",type)
        //如果没有相关数据
        //获取字典
        $.ajax({
            url: process.env.VUE_APP_API_URL+"/dictionary/query",
            async:false,
            type:"GET",
            success:function(data){
              console.log("return data", data)
              EnumUtility.prototype.Set(data.list)
            }
        })
    }
    return JSON.parse(sessionStorage.getItem("menus_"+type)) || [];
};

EnumUtility.prototype.Gets = function (types) {
    if (!types) return JSON.parse(sessionStorage.getItem("menus"));
    var data = [];
    types.split(",").forEach(function (item) {
        data.push(JSON.parse(sessionStorage.getItem("menus"))[item]);
    });
    return data;
};

EnumUtility.prototype.GetTitle = function (type, value) {
    var data = this.Get(type);
    var result = value;
    data.forEach(function (item) {
        if (item.value == value) {
            result = item.description;
            return result;
        }
    });
    return result;
};