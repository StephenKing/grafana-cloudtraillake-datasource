import { E2ESelectors } from '@grafana/e2e-selectors';

export const Components = {
  ConfigEditor: {
    region: {
      input: 'Region',
      wrapper: 'data-testid onloadregion',
    },
    EventDataStore: {
      input: 'Event Data Store',
      wrapper: 'data-testid edsId',
    },
  },
  QueryEditor: {
    CodeEditor: {
      container: 'Code editor container',
    },
    TableView: {
      input: 'toggle-table-view',
    },
  },
  RefreshPicker: {
    runButton: 'RefreshPicker run button',
  },
};

export const selectors: { components: E2ESelectors<typeof Components> } = {
  components: Components,
};
