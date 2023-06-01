import * as item from "../../adapters/item/Item.js";

const container = document.getElementById('table-container');
const data = item.getRows();

data.forEach(rowData => {
    const rowDiv = document.createElement('div');
    rowDiv.setAttribute("class", "table-row");
    Object.keys(rowData).forEach(key => {
        const rowCol = document.createElement('div');
        rowCol.setAttribute("class", "table-row-column");
        const value = rowData[key];
        rowCol.textContent = value;
        rowDiv.append(rowCol);
    });

    const rowCol = document.createElement('div');
    rowCol.setAttribute("class", "table-row-control-column");
    rowDiv.append(rowCol);
    const editButton = document.createElement('a');
    editButton.setAttribute("class", "row-button");
    editButton.setAttribute("href", "./CardItemPage.html?id=" + rowData.id);
    const editIcon = document.createElement('div');
    editIcon.setAttribute("class", "edit-icon");
    editButton.append(editIcon);
    rowCol.append(editButton);
    const deleteButton = document.createElement('div');
    deleteButton.setAttribute("class", "row-delete-button");

    deleteButton.onclick = function () {
        item.deleteRow(rowData.id);
        rowDiv.parentNode.removeChild(rowDiv);
    }

    const deleteIcon = document.createElement('div');
    deleteIcon.setAttribute("class", "delete-icon");
    deleteButton.append(deleteIcon);
    rowCol.append(deleteButton);

    container.append(rowDiv);
});