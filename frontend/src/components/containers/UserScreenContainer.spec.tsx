import { afterEach, beforeEach, describe, it } from 'vitest';
import { RenderResult, render } from '@testing-library/react';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import UserScreenContainer from './UserScreenContainer';

describe('UserScreenContainer', () => {
	let component: RenderResult;
	const queryClient = new QueryClient();

	beforeEach(() => {
		component = render(
			<QueryClientProvider client={queryClient}>
				<UserScreenContainer />
			</QueryClientProvider>
		);
	});

	afterEach(() => {
		component.unmount();
	});

	it('Should render successfully', () => {
		component.getByText('Short your favourite links!');
		component.getByTestId('url-form');
		component.getByLabelText('link-counter');
	});
});
