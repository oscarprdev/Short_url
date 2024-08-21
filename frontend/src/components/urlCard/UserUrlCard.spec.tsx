import { mockedUrl } from '../../../tests/mocks/url-response';
import UserUrlCard from './UserUrlCard';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { RenderResult, render } from '@testing-library/react';
import { afterEach, beforeEach, describe, expect, it } from 'vitest';

describe('UserUrlCard', () => {
	let component: RenderResult;
	const queryClient = new QueryClient();
	const cardsRef: React.MutableRefObject<HTMLLIElement[]> = { current: [] };

	beforeEach(() => {
		component = render(
			<QueryClientProvider client={queryClient}>
				<UserUrlCard url={mockedUrl} cardsRef={cardsRef} isRowsLayout={false} />
			</QueryClientProvider>
		);
	});

	afterEach(() => {
		component.unmount();
	});

	it('Should render successfully', () => {
		component.getByRole('listitem');

		const link = component.getByRole('link');

		expect(link.getAttribute('href')).toBe(mockedUrl.originalUrl);
		expect(link.textContent).toBe(mockedUrl.shortUrl);

		component.getByDisplayValue(mockedUrl.titleUrl);

		component.getByTestId('icon-copy');
		component.getByTestId('icon-pointer');
		component.getByTestId('icon-calendar');
	});
});