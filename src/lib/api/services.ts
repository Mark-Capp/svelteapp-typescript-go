import type { Product, OrderData, Order, Result, Film } from '../type/entities';
const apiBase = "http://127.0.0.1:8080/api";

const api = {
    productsPath: `${apiBase}/products`,
    orderPath: `${apiBase}/orders`,
    filmsPath: `${apiBase}/films`
};

export const loadProducts = async (): Promise<Product[]> => {
    try {
        return await (await fetch(api.productsPath)).json();
    } catch {
        return [];
    }
};

export const loadFilms = async (): Promise<Film[]> => {
    try {
        return await (await fetch(api.filmsPath)).json();
    } catch {
        return [];
    }
};

export const storeOrder = async (order: Order): Promise<Result> => {
    const orderData: OrderData = {
        lines: [...order.orderLines].map((ol) => ({
            productId: ol.product.id,
            productName: ol.product.name,
            quantity: ol.quantity
        }))
    };

    try {
        const result: Result = await (
            await fetch(api.orderPath, {
                method: 'POST',
                headers: {
                    'content-type': 'application/json; charset=UTF-8'
                },
                body: JSON.stringify(orderData)
            })
        ).json();
        return result;
    } catch {
        return { id: -1 };
    }
};
