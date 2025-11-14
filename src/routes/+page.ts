import { getItems } from "$lib/api/services";

export const load = async() => {
    const items = await getItems();
    return {items };
}