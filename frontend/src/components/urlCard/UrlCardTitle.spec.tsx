import { mockedUrl } from '../../../tests/mocks/url-response';
import UrlCardTitle from './UrlCardTitle';
import { QueryClient, QueryClientProvider } from '@tanstack/react-query';
import { RenderResult, fireEvent, render } from '@testing-library/react';
import { afterEach, beforeEach, describe, expect, it } from 'vitest';

describe('UrlCardTitle', () => {
	let component: RenderResult;
	const queryClient = new QueryClient();

	beforeEach(() => {
		component = render(
			<QueryClientProvider client={queryClient}>
				<UrlCardTitle title={mockedUrl.titleUrl} urlId={mockedUrl.id} />
			</QueryClientProvider>
		);
	});

	afterEach(() => {
		component.unmount();
	});

	it('Should render successfully', () => {
		component.getByDisplayValue(mockedUrl.titleUrl);
	});

	it('Should change the value properly onChange event', async () => {
		const input = component.getByDisplayValue(mockedUrl.titleUrl);

		fireEvent.change(input, { target: { value: 123 } });

		expect(input.getAttribute('value')).toBe('123');
	});
});
