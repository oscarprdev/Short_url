import { mockedUrl } from '../../../tests/mocks/url-response';
import UrlCard from './UrlCard';
import { RenderResult, render } from '@testing-library/react';
import { afterEach, beforeEach, describe, expect, it } from 'vitest';

describe('UrlCard', () => {
	let component: RenderResult;

	beforeEach(() => {
		component = render(<UrlCard url={mockedUrl} />);
	});

	afterEach(() => {
		component.unmount();
	});

	it('Should render successfully', () => {
		component.getByRole('listitem');
		const link = component.getByRole('link');

		expect(link.getAttribute('href')).toBe(mockedUrl.originalUrl);
		expect(link.textContent).toBe(mockedUrl.shortUrl);
	});
});
