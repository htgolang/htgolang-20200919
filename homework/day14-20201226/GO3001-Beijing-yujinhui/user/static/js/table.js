// 修改列为可编辑
function trContentStatus(trEle, options) {
    /*
        editable    boolean     是否可编辑
        normal      boolean     正常形式（不可编辑）
        editClass   string      编辑类
    */
    const {editable, normal, editClass} = options;
    let rwCols = trEle.getElementsByClassName('rw');
    let rdCols = trEle.getElementsByClassName('rd');

    // 正常表格
    if (normal) {
        // 移除编辑样式
        trEle.classList.remove(editClass);
        for (let i = 0; i < rwCols.length; i++) {
            rwCols[i].setAttribute("contenteditable", false);
        }
    }

    // 可编辑表格
    if (editable) {
        // 添加编辑样式
        trEle.classList.add(editClass);
        for (let i = 0; i < rwCols.length; i++) {
            rwCols[i].setAttribute("contenteditable", true);
        }
    }
}

// 创建按钮
function trContentBtn(trEle, options) {
    /*
        value       string      按钮显示文字
        btnClassArr Array       按钮样式
        clickEvent  function    点击事件
    */
    const {value, btnClassArr, clickEvent} = options;

    // 创建按钮
    const Btn = document.createElement("button");
    Btn.innerText = value;

    // 添加样式
    for (item of btnClassArr) {
        Btn.classList.add(item);
    }

    // 按钮点击事件
    Btn.addEventListener("click", (btn) => {
        clickEvent(btn)
    }, false)

    // 挂载到Dom节点
    trEle.children[trEle.children.length - 1].appendChild(Btn);
}