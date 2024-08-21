import ShortUrlForm from './ShortUrlForm';
import { RenderResult, render } from '@testing-library/react';
import { afterEach, beforeEach, describe, it, vi } from 'vitest';

describe('ShortUrlForm', () => {
	let component: RenderResult;
	const addUrl = vi.fn();

	beforeEach(() => {
		component = render(<ShortUrlForm addUrl={addUrl} userId="1" />);
	});

	afterEach(() => {
		component.unmount();
	});

	it('Should render successfully', async () => {
		component.getByTestId('url-form');
		component.getByTestId('submit-btn');
		component.getByPlaceholderText('Enter your link here...');
	});
});
