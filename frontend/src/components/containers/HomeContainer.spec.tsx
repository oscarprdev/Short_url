import { afterEach, beforeEach, describe, it } from 'vitest';
import { RenderResult, render } from '@testing-library/react';
import HomeContainer from './HomeContainer';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';

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
		component.getByText('Short - it');
		component.getByText('The simplest URL Shortner you were waiting for.');
		component.getByTestId('url-form');
	});
});
