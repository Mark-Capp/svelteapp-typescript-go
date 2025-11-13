import type { PageLoad } from '../$types';
import { loadFilms } from '$lib/api/services';

export const load: PageLoad = async () => {
    // load product data from backend service before mounting Product component
    return {
        films: await loadFilms()
    };
};
