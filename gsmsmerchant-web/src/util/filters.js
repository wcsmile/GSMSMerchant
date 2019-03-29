import Vue from "vue"
import { EnumUtility } from './enum'

Vue.filter('EnumFilter', (value, enumType) => {
    if (value !== '') {
        let enumUtility = new EnumUtility()
        let enumMap = enumUtility.Get(enumType)
        let result = value
        enumMap.forEach(item => {
            if (item.value === value) {
                result = item.name
            }
        })
        return result
    } else {
        return '-'
    }
})


Vue.filter('DateFilter', (value, format) => {
    let res
    if (value === '') {
        return '-'
    } else {
        res = DateConvert(format, value)
        return res
    }
})

Vue.filter('StringFilter', value => {
    if (value === '') {
        return '---'
    } else {
        return value
    }
})

Vue.filter('EllipsisFilter', (value, number) => {
    if (value) {
        if (value.length <= number) {
            return value
        } else {
            let subval = value.slice(0, number - 1) + '...'
            return subval
        }
    } else {
        return '---'
    }
})

Vue.filter('FinalFilter', (value, number) => {

    if (value) {
        if (number < 9) {
            return value
        } else {
            let subval = '...' + value.slice(value.length - number, value.length - 1)
            return subval
        }
    } else {
        return '-'
    }
})

Vue.filter('RemarkFilter', (value) => {
    if (value === '') {
        return '-'
    } else {
        return value
    }
})

Vue.filter('AmountFilter', (value) => {
    if (value == null || value == undefined || value == "") {
        return 0
    }

    if (typeof(value) != "string") {
        value = value.toString()
    }

    if (value.indexOf('-') === -1) {
        return Common.dealwithNum(value)
    } else {
        let result = value.slice(1, value.length)
        return '-' + Common.dealwithNum(result)
    }
})


Vue.filter('BusinesstTypeFilter', (value, list) => {
    let result = value
    list.forEach(item => {
        if (result === item.business_type) {
            result = item.business_name
        }
    })
    return result
})

Vue.filter('ChannelTypeFilter', (value, list)=>{
    let result = value
    list.forEach(item => {
        if (result === item.value) {
            result = item.name
        }
    })
    return result
})

Vue.filter('BranchFilter', (value, list) => {
    let result = value
    list.forEach(item => {
        if (result === item.branch_id) {
            result = item.branch_name
        }
    })
    return result
})
