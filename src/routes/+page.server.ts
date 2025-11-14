import { addItem, addTag } from '$lib/api/services';
import type { Actions, Cookies } from '@sveltejs/kit';


export const actions = {
	addItem: async ({ request } : {request: Request}) => {
		const form = await request.formData();
        const itemValue = String(form.get('item'));
        await addItem(itemValue);
	},
    addTag: async ({ request } : {request: Request}) => {
        const form = await request.formData();
        const tagValue = String(form.get('tag'));
        console.log("Adding tag: ", tagValue);
        await addTag(tagValue);
    }
} as Actions;