import Footer from './Footer';
import { RenderResult, render } from '@testing-library/react';
import { afterEach, beforeEach, describe, expect, it } from 'vitest';

describe('Footer', () => {
	let component: RenderResult;

	beforeEach(() => {
		component = render(<Footer />);
	});

	afterEach(() => {
		component.unmount();
	});

	it('Should render successfully', async () => {
		component.getByTestId('heart-icon');

		const link = component.getByRole('link');

		expect(link.textContent).toBe('@oscarprdev');
		expect(link.getAttribute('href')).toBe('https://github.com/oscarprdev');
	});
});
