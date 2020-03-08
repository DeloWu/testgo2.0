// 根据索引移除
function removeInArrayByIndex(array, removeIndex){
    let tempArray = [];
    for(let i=0;i<array.length;i++){
        if(i != removeIndex){
            tempArray.push(array[i]);
        }
    }
    return tempArray
}

//专门针对表格数据，将value转换为对应的数据类型
//[{"keyName":"foo", "valueType":"int", "value":"1.0"}, {, , }...] ==> [{"keyName":"foo", "valueType":"int", "value":"1"}, {, , }...]
function transformValue(arr) {
    for (let item of arr){
        window.console.log(item);

    }
}
export {removeInArrayByIndex, transformValue}