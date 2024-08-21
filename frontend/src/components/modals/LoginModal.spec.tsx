import { API_URL } from '../../constants/apiUrl';
import LoginModal from './LoginModal';
import { RenderResult, fireEvent, render, waitFor } from '@testing-library/react';
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';

describe('LoginModal', () => {
	let component: RenderResult;

	beforeEach(() => {
		component = render(<LoginModal />);

		Object.defineProperty(window, 'location', {
			writable: true,
			value: { assign: vi.fn() },
		});
	});

	afterEach(() => {
		component.unmount();
	});

	it('Should update location when clicks on default button', async () => {
		const buttons = component.getAllByRole('link');

		const defaultButton = buttons.find(btn => btn.innerText === 'Default user');

		if (defaultButton) {
			fireEvent.click(defaultButton);

			await waitFor(() => {
				expect(window.location.href).toBe('https://opr-short-url.vercel.app/user/116176187754032784002');
			});
		}
	});

	it('Should update location when clicks on google button', async () => {
		const buttons = component.getAllByRole('link');

		const defaultButton = buttons.find(btn => btn.innerText === 'Google');

		if (defaultButton) {
			fireEvent.click(defaultButton);

			await waitFor(() => {
				expect(window.location.href).toBe(`${API_URL}/auth?provider=google`);
			});
		}
	});

	it('Should update location when clicks on github button', async () => {
		const buttons = component.getAllByRole('link');

		const defaultButton = buttons.find(btn => btn.innerText === 'Github');

		if (defaultButton) {
			fireEvent.click(defaultButton);

			await waitFor(() => {
				expect(window.location.href).toBe(`${API_URL}/auth?provider=github`);
			});
		}
	});
});
