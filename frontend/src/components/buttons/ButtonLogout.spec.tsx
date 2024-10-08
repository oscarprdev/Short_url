import { API_URL } from '../../constants/apiUrl';
import ButtonLogout from './ButtonLogout';
import { RenderResult, fireEvent, render, waitFor } from '@testing-library/react';
import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';

describe('ButtonLogout', () => {
	let component: RenderResult;

	beforeEach(() => {
		component = render(<ButtonLogout />);

		Object.defineProperty(window, 'location', {
			writable: true,
			value: { assign: vi.fn() },
		});
	});

	afterEach(() => {
		component.unmount();
	});

	it('Should render successfully', async () => {
		component.getByTestId('logout-btn');
	});

	it('Should update location when clicks on button', async () => {
		const button = component.getByText('Logout');

		fireEvent.click(button);

		await waitFor(() => {
			expect(window.location.href).toBe(`${API_URL}/auth/logout`);
		});
	});
});
