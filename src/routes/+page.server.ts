import { addItem } from '$lib/api/services';
import type { Actions, Cookies } from '@sveltejs/kit';


export const actions = {
	default: async ({ cookies, request } : {cookies: Cookies, request: Request}) => {
		const form = await request.formData();
        const itemValue = String(form.get('item'));
        // Call the addItem function to add the item
        await addItem(String(form.get('item')));
	}
} as Actions;