import type { Product, OrderData, Order, Result } from '../type/entities';

import { dev } from '$app/environment';
import mockProducts from '$lib/data/mockData.json';

const api = {
    loadProductsPath: 'http://127.0.0.1:8080/api/products',
    storeOrderPath: 'http://127.0.0.1:8080/api/orders'
};

export const loadProducts = async (): Promise<Product[]> => {

    try {
        return await (await fetch(api.loadProductsPath)).json();
    } catch {
        return [];
    }
};

export const storeOrder = async (order: Order): Promise<Result> => {

    const orderData: OrderData = {
        lines: [...order.orderLines.currents()].map((ol) => ({
            productId: ol.product.id,
            productName: ol.product.name,
            quantity: ol.quantity
        }))
    };

    try {
        const result: Result = await (
            await fetch(api.storeOrderPath, {
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
