import HomeContainer from './HomeContainer';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { RenderResult, render } from '@testing-library/react';
import { afterEach, beforeEach, describe, it } from 'vitest';

describe('HomeContainer', () => {
	let component: RenderResult;
	const queryClient = new QueryClient();

	beforeEach(() => {
		component = render(
			<QueryClientProvider client={queryClient}>
				<HomeContainer />
			</QueryClientProvider>
		);
	});

	afterEach(() => {
		component.unmount();
	});

	it('Should render successfully', () => {
		component.getByText('OPlink');
		component.getByText('The simplest URL Shortner app you were waiting for.');
		component.getByTestId('url-form');
	});
});
