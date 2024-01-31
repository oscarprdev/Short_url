import { describe, expect, it } from 'vitest';
import { render } from '@testing-library/react';
import AuthInput from './AuthInput';
import { mockedUrl } from '../../tests/mocks/url-response';

describe('AuthInput', () => {
	it('Should render successfully', async () => {
		const component = render(
			<AuthInput
				children={<p>Hello</p>}
				contrast={true}
				href={mockedUrl.originalUrl}
			/>
		);

		const link = component.getByRole('link');

		expect(link.getAttribute('href')).toBe(mockedUrl.originalUrl);

		component.getByText('Hello');

		expect(link.className).toContain('bg-[var(--contrast)] hover:bg-[var(--contrast-hover)]');

		component.unmount();
	});

	it('Should render proper styles when constrast is false', async () => {
		const component = render(
			<AuthInput
				children={<p>Hello</p>}
				contrast={false}
				href={mockedUrl.originalUrl}
			/>
		);

		const link = component.getByRole('link');

		expect(link.className).toContain('bg-primary hover:bg-primary/90');

		component.unmount();
	});
});
