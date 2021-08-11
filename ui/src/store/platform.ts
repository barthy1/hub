import { types, Instance } from 'mobx-state-tree';

export const Platform = types.model('Platforms', {
  id: types.number,
  name: types.identifier
});

export const PlatformStore = types
  .model({
    items: types.optional(types.map(Platform), {})
  })
  .actions((self) => ({
    add(item: IPlatform): void {
      self.items.put(item);
    }
  }))
  .views((self) => ({
    get values() {
      return Array.from(self.items.values());
    }
  }));

export type IPlatform = Instance<typeof Platform>;
export type IplatformStore = Instance<typeof PlatformStore>;
