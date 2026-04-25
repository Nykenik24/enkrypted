import type { Actions } from './$types';

export const actions: Actions = {
    default: async ({ request }) => {
        const formData = await request.formData();
        const username = formData.get('username');

        console.log('Login attempt for:', username);

        // TODO: authentication logic here

        return {
            success: true,
            username
        };
    }
};
