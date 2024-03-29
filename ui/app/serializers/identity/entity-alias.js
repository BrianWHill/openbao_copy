/**
 * Copyright (c) HashiCorp, Inc.
 * SPDX-License-Identifier: MPL-2.0
 */

import IdentitySerializer from './_base';
export default IdentitySerializer.extend({
  extractLazyPaginatedData(payload) {
    return payload.data.keys.map((key) => {
      const model = payload.data.key_info[key];
      model.id = key;
      if (payload.backend) {
        model.backend = payload.backend;
      }
      return model;
    });
  },
});
