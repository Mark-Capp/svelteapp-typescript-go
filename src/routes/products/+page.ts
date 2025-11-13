import type { PageLoad } from './$types';
import { loadProducts } from '$lib/api/services';

export const load: PageLoad = async () => {
    // load product data from backend service before mounting Product component
    var data= {
        products: await loadProducts()
    };
    console.log(data);
    return data;
};
