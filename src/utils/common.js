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

export {removeInArrayByIndex}