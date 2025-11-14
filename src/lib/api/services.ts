import type { Item } from "$lib/type/entities";

const apiBase = "http://127.0.0.1:8080/api";

const api = {
    itemsPath: `${apiBase}/items`
};

export const addItem = async (itemName: string): Promise<void> => {
    try {
        await fetch(api.itemsPath, {
            method: 'POST',
            headers: {
                'content-type': 'application/json; charset=UTF-8'
            },
            body: JSON.stringify({ name: itemName })
        });    
    } catch {
        console.error("Failed to add item");
    }
};

export const getItems = async (): Promise<Item[]> => {
    try {
        return await (await fetch(api.itemsPath)).json();
    } catch {
        return []
    }
};

