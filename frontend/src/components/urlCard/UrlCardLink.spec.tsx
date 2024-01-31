import { afterEach, beforeEach, describe, expect, it, vi } from 'vitest';
import { RenderResult, fireEvent, render } from '@testing-library/react';
import UrlCardLink from './UrlCardLink';

describe('UrlCardLink', () => {
	let component: RenderResult;
	const originalUrl = 'http:/short&go.dev/long/url';
	const shortUrl = 'http:/short&go.dev/short/url';
	const onUrlClick = vi.fn();

	beforeEach(() => {
		component = render(
			<UrlCardLink
				originalUrl={originalUrl}
				shortUrl={shortUrl}
				onUrlClick={onUrlClick}
			/>
		);
	});

	afterEach(() => {
		component.unmount();
	});

	it('Should render successfully', async () => {
		component.getByLabelText('url-link');
		const link = component.getByRole('link');

		expect(link.textContent).toBe(shortUrl);
		expect(link.getAttribute('href')).toBe(originalUrl);
	});

	it('Should trigger click event', async () => {
		const link = component.getByRole('link');

		fireEvent.click(link);

		expect(onUrlClick).toHaveBeenCalled();
	});
});
